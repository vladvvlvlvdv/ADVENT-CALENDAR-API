package repository

type (
	Setting struct {
		ID          uint `json:"id"`
		Month       int  `json:"month" gorm:"not null"`
		ShowAllDays bool `json:"showAllDays" gorm:"not null;default:false"`
	}
)

var SettingService = new(Setting)

func (s *Setting) Get() (Setting, error) {
	if err := DB.Model(s).First(&s); err != nil {
		return Setting{}, err.Error
	}

	return *s, nil
}
