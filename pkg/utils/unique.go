package utils

import (
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

/*
Генерирует уникальное название файла
*/
func GenerateUniqueFilename(file *multipart.FileHeader) string {
	extFile := strings.ToLower(filepath.Ext(file.Filename))

	return uuid.New().String() + extFile
}

func GenerateCode() string {
	rand.Seed(time.Now().UnixNano())
	const chars = "0123456789"
	length := 6

	var code string
	for i := 0; i < length; i++ {
		code += string(chars[rand.Intn(len(chars))])
	}

	return code
}
