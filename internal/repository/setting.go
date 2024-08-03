package repository

type (
	SettingDTO struct {
		Month       int  `form:"month" validate:"required,min=1,max=12"`
		ShowAllDays bool `form:"showAllDays"`
	}
	Setting struct {
		ID          uint `json:"id"`
		Month       int  `json:"month" gorm:"not null"`
		ShowAllDays bool `json:"showAllDays" gorm:"not null;default:false"`
	}
)

var SettingService = new(Setting)

func (s Setting) Get() (Setting, error) {
	if err := DB.Model(s).First(&s); err != nil {
		return Setting{}, err.Error
	}

	return s, nil
}

func (s Setting) Update(toUpdate Setting) error {
	return DB.Model(&s).Updates(&toUpdate).Error
}
