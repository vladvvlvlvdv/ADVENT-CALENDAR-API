package repository

import (
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

func (a Attachment) DeleteMany(ids []uint) error {
	if err := DB.Model(a).Where("id IN(?)", ids).Delete(Attachment{}).Error; err != nil {
		return errors.New("Ошибка при удалении вложений")
	}

	return nil
}

func (a *Attachment) BeforeDelete(tx *gorm.DB) (err error) {

	os.Remove(fmt.Sprintf("./%s", a.URL))

	return
}
