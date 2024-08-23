package repository

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

type (
	Project struct {
		ID          uint   `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Link        string `json:"link"`
		Preview     string `json:"preview"`
	}

	ProjectDTO struct {
		Title       string `form:"title" validate:"required,min=5"`
		Description string `form:"description" validate:"required,min=5"`
		Link        string `form:"link" validate:"required,url"`
	}
)

var ProjectService = new(Project)

func (p *Project) Create(project Project) error {
	return DB.Create(&Project{
		Title:       project.Title,
		Description: project.Description,
		Preview:     project.Preview,
		Link:        project.Link,
	}).Error
}

func (p *Project) GetAll(where Project) ([]Project, error) {
	var projects []Project
	return projects, DB.Where(where).Find(&projects).Error
}

func (p *Project) Get(project Project) (Project, error) {
	return project, DB.Where(project).First(&project).Error
}

func (p *Project) Update() error {
	return DB.Where("id", p.ID).Updates(p).Error
}

func (p *Project) Delete() error {
	return DB.Delete(p).Error
}

func (p *Project) BeforeDelete(tx *gorm.DB) (err error) {
	os.Remove(fmt.Sprintf("./%s", p.Preview))

	return
}
