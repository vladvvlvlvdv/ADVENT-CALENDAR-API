package utils

import (
	"advent-calendar/pkg/validators"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

type (
	Upload struct {
		FileKey  string
		FileType string
		MaxCount int
		Require  bool
	}

	File struct {
		OriginalName string
		Destination  string
		FileType     string
		Size         int64
	}
)

func UploadFiles(uploads []Upload) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		filesMap := make(map[string][]File)

		data, err := c.MultipartForm()
		if err != nil {
			c.Locals("files", filesMap)
			return c.Next()
		}

		for _, upl := range uploads {
			files := data.File[upl.FileKey]

			if len(files) == 0 && upl.Require {
				return fiber.NewError(400, fmt.Sprintf("Прикрепите файл для: %v", upl.FileKey))
			}

			if upl.MaxCount != 0 && len(files) > upl.MaxCount {
				return fiber.NewError(400, fmt.Sprintf("Максимальное кол-во файлов для %v: %v", upl.FileKey, upl.MaxCount))
			}

			if upl.FileType != "" {
				for _, file := range files {
					if upl.FileType == "" {
						continue
					}
					err := validators.CheckFileExtension(upl.FileType, file)
					if err != nil {
						return fiber.NewError(400, err.Error())
					}
				}
			}
		}

		for _, upl := range uploads {
			files := data.File[upl.FileKey]

			for _, file := range files {
				destination := fmt.Sprintf("./public/%s/%s", upl.FileKey, GenerateUniqueFilename(file))

				if err := c.SaveFile(file, destination); err != nil {
					return fiber.NewError(500, "Ошибка сохранения файла")
				}

				filesMap[upl.FileKey] = append(filesMap[upl.FileKey], File{
					Destination:  destination[1:],
					OriginalName: file.Filename,
					FileType:     validators.GetFileType(file.Filename),
					Size:         file.Size,
				})
			}
		}

		c.Locals("files", filesMap)
		return c.Next()
	}
}

func DeleteFiles(files map[string][]File) error {
	for _, fileKey := range files {
		for _, file := range fileKey {
			err := os.Remove(fmt.Sprintf("./%s", file.Destination))
			if err != nil {
				continue
			}
		}
	}
	return nil
}
