package repository

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type (
	LoginDTO struct {
		Email    string `json:"email" form:"email" validate:"required,min=5,email"`
		Password string `json:"password" form:"password"`
	}

	ConfirmUser struct {
		Code  string `json:"code" form:"code" validate:"required,len=6"`
		Email string `json:"email" form:"email" validate:"required,min=5,email"`
	}

	SubscribeDTO struct {
		Email      string `form:"email" validate:"required,min=5,email"`
		Nickname   string `form:"nickname" validate:"required,min=5"`
		TgUsername string `form:"tgUsername" validate:"required,min=5"`
		IsConfirm  bool   `form:"isConfirm" validate:"required,boolean"`
	}

	UnSubscribeDTO struct {
		Email string `json:"email" validate:"required,min=5,email"`
	}

	User struct {
		ID           uint   `json:"id"`
		Email        string `json:"email" gorm:"unique;not null"`
		Password     string `json:"-" gorm:"not null;type:varchar(255)"`
		Role         string `json:"role" gorm:"not null; default:user"`
		RefreshToken string `json:"refreshToken" gorm:"not null"`
		Code         string `json:"code"`
	}

	Subscribe struct {
		ID         uint   `json:"id"`
		Email      string `json:"email" gorm:"unique;not null"`
		Nickname   string `json:"nickname"`
		TgUsername string `json:"tgUsername"`
		Days       []Day  `gorm:"many2many:days_views;" json:"-"`
	}
)

var UserService = new(User)

func (u User) Get(where User) (User, error) {
	var user User
	err := DB.Where(&where).First(&user).Error
	return user, err
}

func (u User) Update(toUpdate User) error {
	return DB.Model(&u).Where(u).Updates(&toUpdate).Error
}

func (u User) Create(newUser User) (User, error) {
	if err := DB.Create(&newUser).Error; err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			if strings.Contains(err.Error(), "users.uni_users_email") {
				return User{}, errors.New("Такая почта уже используется")
			}
		}
	}
	return newUser, nil
}

func (u User) GetAll(where User) ([]User, error) {
	var users []User

	if err := DB.Where(&where).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u User) Subscribe(s *SubscribeDTO) error {
	if err := DB.Where("email =?", s.Email).First(&Subscribe{}).Error; err == nil {
		return errors.New("На этот email подписка уже есть")
	}
	return nil
}

func (u User) GetAllSubscribes() ([]Subscribe, error) {
	var subs []Subscribe

	if err := DB.Preload("Days", func(db *gorm.DB) *gorm.DB {
		return db.Select("id")
	}).Find(&subs).Error; err != nil {
		return nil, err
	}

	return subs, nil
}

func (u User) GetSubscriber(s Subscribe) (Subscribe, error) {
	var sub Subscribe

	if err := DB.Where(&s).First(&sub).Error; err != nil {
		return Subscribe{}, err
	}

	return sub, nil
}

func (u User) UnSubscribe(email string) error {
	return DB.Where("email = ?", email).Delete(&Subscribe{}).Error
}
