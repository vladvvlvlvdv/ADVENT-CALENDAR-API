package repository

import (
	"advent-calendar/internal/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var DB *gorm.DB

func LoadDatabase() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Config.DB_USER,
		config.Config.DB_PASSWORD,
		config.Config.DB_HOST,
		config.Config.DB_PORT,
		config.Config.DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных")
	}
}

func AutoMigrate() {
	if err := DB.AutoMigrate(
		&Day{},
		&Attachment{},
		&Setting{},
	); err != nil {
		log.Fatal("Ошибка миграции таблиц")
	}
}

func RenderDatabase() {
	DB.Where("id = 1").FirstOrCreate(&Setting{
		Month: 12,
	})
}
