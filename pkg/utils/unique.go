package utils

import (
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

/*
Генерирует уникальное название файла
*/
func GenerateUniqueFilename(file *multipart.FileHeader) string {
	extFile := strings.ToLower(filepath.Ext(file.Filename))

	return uuid.New().String() + extFile
}
