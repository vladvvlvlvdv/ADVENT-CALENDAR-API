package repository

import "gorm.io/gorm"

type (
	Day struct {
		Model
		Title       string       `json:"title"`
		Description string       `json:"description"`
		Attachments []Attachment `json:"attachments"`
	}

	DayDTO struct {
		Title       string       `json:"title" form:"title"`
		Description string       `json:"description" form:"description"`
		Attachments []Attachment `json:"-" form:"-"`
	}
)

var DayService = new(Day)

func (d *Day) Create(day DayDTO) error {
	return DB.Model(d).Create(&Day{
		Title:       day.Title,
		Description: day.Description,
		Attachments: day.Attachments,
	}).Error
}

/* Scopes */

func (d *Day) GetAttachments() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("Attachments")
	}
}
