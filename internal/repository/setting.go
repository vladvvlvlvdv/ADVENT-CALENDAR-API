package repository

type (
	Setting struct {
		Model
		Month int `json:"month" gorm:"not null"`
	}
)
