package repository

type (
	Click struct {
		IP string `json:"ip" gorm:"primaryKey;not null;unique"`
	}
)

var ClickService = new(Click)

func (c Click) Create(newClick Click) error {
	return DB.Where(newClick).FirstOrCreate(&newClick).Error
}

func (c Click) Count() (int64, error) {
	var count int64

	if err := DB.Model(&c).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
