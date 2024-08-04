package validators

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"
)

var FilesExtensions = map[string][]string{
	"image": {".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".tif", ".raw", ".svg", ".webp", ".heic", ".heif", ".ico", ".jfif", ".pjpeg", ".pjp", ".avif"},
	"video": {".mp4", ".mov", ".avi", ".mkv", ".wmv", ".flv", ".swf", ".mpeg", ".mpg", ".3gp", ".webm", ".vob", ".ogv", ".mts", ".m2ts", ".mxf", ".m4v", ".divx", ".xvid", ".rm", ".rmvb", ".asf", ".dv"},
	"audio": {".mp3", ".wav", ".ogg", ".aac", ".flac", ".wma", ".m4a", ".aiff", ".ape", ".alac", ".amr", ".opus", ".ra", ".mid", ".midi", ".mka", ".dsf", ".dff", ".wv"},
}

func CheckFileExtension(checkType string, file *multipart.FileHeader) error {
	extFile := strings.ToLower(filepath.Ext(file.Filename))

	for _, ext := range FilesExtensions[checkType] {
		if extFile == ext {
			return nil
		}
	}

	return errors.New("Недопустимый формат файла")
}

/*
filetype 1 - остальные файлы, 2 - фото, 3 - видео, 4 - аудио
*/
func GetFileType(filename string) string {
	extFile := strings.ToLower(filepath.Ext(filename))

	for key, ext := range FilesExtensions {
		for _, value := range ext {
			if extFile == value {
				return key
			}
		}
	}

	return ""
}
