package repository

type (
	LoginDTO struct {
		Email    string `json:"email" form:"email" validate:"required,min=5,email"`
		Password string `json:"password" form:"password" validate:"required,min=5"`
	}

	ConfirmUser struct {
		Code  string `json:"code" form:"code" validate:"required,len=6"`
		Email string `json:"email" form:"email" validate:"required,min=5,email"`
	}

	User struct {
		ID           uint   `json:"id"`
		Email        string `json:"email" gorm:"unique;not null"`
		Password     string `json:"-" gorm:"not null;type:varchar(255)"`
		Role         string `json:"role" gorm:"not null; default:user"`
		RefreshToken string `json:"refreshToken" gorm:"not null"`
		Code         string `json:"code"`
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
	err := DB.Create(&newUser).Error
	return newUser, err
}
