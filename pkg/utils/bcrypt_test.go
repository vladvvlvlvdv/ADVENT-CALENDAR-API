package utils

import "testing"

func TestBcryptPassword(t *testing.T) {
	testPassword := "test_password"

	hash, err := HashPassword(testPassword)
	if err != nil {
		t.Fatalf("Ошибка при хешировании пароля: %v", err)
	}

	ok := CheckPasswordHash(testPassword, hash)
	if !ok {
		t.Error("Пароль не верно хеширован")
	}

	t.Log("Тест bcrypt успешно завершен")
}
