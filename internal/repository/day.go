package repository

import (
	"advent-calendar/pkg/utils"
)

type (
	Day struct {
		ID          uint         `json:"id"`
		Title       string       `json:"title"`
		Description string       `json:"description"`
		IsLongRead  bool         `json:"isLongRead"`
		Attachments []Attachment `json:"attachments,omitempty"`
		Users       []User       `gorm:"many2many:days_views;"`
	}

	DayDTO struct {
		Title       string `json:"title" form:"title" validate:"required,min=5"`
		Description string `json:"description" form:"description" validate:"required,min=5"`
	}

	DayUPD struct {
		Title         string `form:"title" validate:"min=5"`
		Description   string `form:"description" validate:"min=5"`
		AttachmentIds []uint `form:"attachmentIds"`
	}

	DayView struct {
		UserID uint `gorm:"primaryKey"`
		DayID  uint `gorm:"primaryKey"`
	}
)

var DayService = new(Day)

func (d Day) Create(day DayDTO, files []utils.File) error {

	if len(files) > 0 {
		for _, file := range files {
			d.Attachments = append(d.Attachments, Attachment{
				Label: file.OriginalName,
				URL:   file.Destination,
				Type:  file.FileType,
			})
		}
	}

	d.Title = day.Title
	d.Description = day.Description

	return DB.Model(&d).Create(&d).Error
}

func (d Day) GetAll(params Params, where Day) ([]Day, error) {
	var days []Day

	query := DB.Model(&d).Where("id <= ?", where.ID).Preload("Attachments")

	if params.Limit > 0 {
		query = query.Limit(params.Limit).Offset((params.Page - 1) * params.Limit)
	}

	if err := query.Find(&days).Error; err != nil {
		return nil, err
	}

	return days, nil
}

func (d *Day) Update(id uint, day DayUPD, files []utils.File) error {
	if err := DB.Model(&d).Where("id = ?", id).Updates(Day{Title: day.Title, Description: day.Description}).Error; err != nil {
		return err
	}

	if err := AttachmentService.CreateMany(files, id); err != nil {
		return err
	}

	return nil
}

func (d Day) Get(where Day) (Day, error) {
	if err := DB.Model(&d).Preload("Attachments").Where(where).First(&d).Error; err != nil {
		return Day{}, err
	}

	return d, nil
}

func (d Day) CreateView(newView DayView) error {
	return DB.Model(DayView{}).Where(newView).FirstOrCreate(&newView).Error
}
