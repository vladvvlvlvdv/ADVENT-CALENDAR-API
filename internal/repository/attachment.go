package repository

import (
	"advent-calendar/pkg/utils"
	"errors"
	"fmt"
	"os"

	"gorm.io/gorm"
)

type (
	Attachment struct {
		ID    uint   `json:"id"`
		Label string `json:"label"`
		URL   string `json:"url"`
		Type  string `json:"type"`
		DayID uint   `json:"-"`
	}
)

var AttachmentService = new(Attachment)

func (a Attachment) DeleteMany(ids []uint) error {
	if len(ids) == 0 {
		return nil
	}

	var attIds []Attachment
	for _, id := range ids {
		attIds = append(attIds, Attachment{ID: id})
	}

	if err := DB.Model(&a).Delete(&attIds).Error; err != nil {
		return errors.New("Ошибка при удалении вложений")
	}

	return nil
}

func (a Attachment) CreateMany(files []utils.File, dayId uint) error {
	if len(files) > 0 {
		var attachments []Attachment

		for _, file := range files {
			attachments = append(attachments, Attachment{
				Label: file.OriginalName,
				URL:   file.Destination,
				Type:  file.FileType,
				DayID: dayId,
			})
		}

		if err := DB.Create(&attachments).Error; err != nil {
			return errors.New("Ошибка при создании вложений")
		}
	}

	return nil
}

func (a *Attachment) BeforeDelete(tx *gorm.DB) (err error) {
	return os.Remove(fmt.Sprintf("./%s", a.URL))
}
