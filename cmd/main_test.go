package main

import (
	"advent-calendar/internal/config"
	"advent-calendar/internal/handler"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestE2EMain(t *testing.T) {
	if err := os.Chdir("../"); err != nil {
		panic(fmt.Sprintf("Не удалось сменить рабочую директорию: %v", err))
	}

	go func() {
		main()
	}()

	time.Sleep(2 * time.Second)

	host := fmt.Sprintf("http://localhost:%s", config.Config.PORT)

	// * Тестирование получения списка дней
	respDays, err := http.Get(fmt.Sprintf("%s/api/days", host))
	require.NoError(t, err)
	if err == nil {
		require.Equal(t, 200, respDays.StatusCode)
	}

	t.Log("Список дней успешно получен")

	// * Тестирование авторизации пользователя
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	writer.WriteField("email", "admin@example.com")
	writer.WriteField("password", "admin")

	err = writer.Close()
	if err != nil {
		t.Fatalf("Ошибка при формировании Login формы: %v", err)
	}

	respLogin, err := http.Post(fmt.Sprintf("%s/api/users/login", host), writer.FormDataContentType(), &buf)
	require.NoError(t, err)
	if err == nil {
		require.Equal(t, 200, respLogin.StatusCode)
	}

	var tokens handler.Tokens
	if respLogin.StatusCode == 200 {

		err := json.NewDecoder(respLogin.Body).Decode(&tokens)
		if err != nil {
			t.Fatalf("Ошибка при парсинге тела ответа: %v", err)
		}

		require.NotEmpty(t, tokens.AccessToken, "Токен авторизации должен быть в ответе")
	}

	t.Log("Авторизация прошла успешно")

	authToken := fmt.Sprintf("Bearer %s", tokens.AccessToken)

	// * Тестирование создания проекта
	var projectBuf bytes.Buffer
	projectWriter := multipart.NewWriter(&projectBuf)

	projectWriter.WriteField("title", "Тестовый заголовок")
	projectWriter.WriteField("description", "Тестовое описание")
	projectWriter.WriteField("link", "https://test.com/link")

	filePath := "./test_assets/test_preview.png"

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Ошибка при открытии файла: %v", err)
	}

	fileWriter, err := projectWriter.CreateFormFile("previews", filepath.Base(filePath))
	if err != nil {
		t.Fatalf("Ошибка при добавлении файла в форму: %v", err)
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		t.Fatalf("Ошибка при копировании файла в форму: %v", err)
	}

	err = projectWriter.Close()
	if err != nil {
		t.Fatalf("Ошибка при формировании Project формы: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/projects", host), &projectBuf)
	if err != nil {
		t.Fatalf("Ошибка при создании запроса: %v", err)
	}

	req.Header.Set("Content-Type", projectWriter.FormDataContentType())
	req.Header.Set("Authorization", authToken)

	client := &http.Client{}
	respProject, err := client.Do(req)
	require.NoError(t, err)

	if err == nil {
		require.Equal(t, 200, respProject.StatusCode)
	}

	t.Log("Проект успешно создан")
}
