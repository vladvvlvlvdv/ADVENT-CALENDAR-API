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
		IsViewed    bool         `json:"isViewed"`
		Attachments []Attachment `json:"attachments,omitempty"`
		Users       []Subscribe  `gorm:"many2many:days_views;" json:"-"`
	}

	DayDTO struct {
		Title       string `json:"title" form:"title" validate:"required,min=5"`
		Description string `json:"description" form:"description" validate:"required,min=5"`
	}

	DayUPD struct {
		Title         string `form:"title" validate:"min=5"`
		Description   string `form:"description" validate:"min=5"`
		IsLongRead    bool   `form:"isLongRead" validate:"boolean"`
		AttachmentIds []uint `form:"attachmentIds"`
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

	subQuery := DB.Table("days_views").
		Select("true").
		Where("days_views.day_id = days.id").
		Where("days_views.subscribe_id = ?", params.SubscribeId).
		Limit(1)

	query := DB.Model(&d).
		Select("days.*, (?) AS is_viewed", subQuery).
		Where("days.id <= ?", where.ID).
		Preload("Attachments")

	if params.Limit > 0 {
		query = query.Limit(params.Limit).Offset((params.Page - 1) * params.Limit)
	}

	if err := query.Find(&days).Error; err != nil {
		return nil, err
	}

	return days, nil
}

func (d *Day) Update(id uint, day DayUPD, files []utils.File) error {
	if err := DB.Model(&d).Where("id = ?", id).
		Update("title", day.Title).
		Update("description", day.Description).
		Update("is_long_read", day.IsLongRead).
		Error; err != nil {
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

func (d Day) CreateView(subId, dayID uint) error {
	return DB.Model(&Subscribe{ID: subId}).Association("Days").Append(&Day{ID: dayID})
}
