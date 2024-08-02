package repository

import (
	"advent-calendar/internal/config"
	"advent-calendar/pkg/utils"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	Params struct {
		Limit int
		Page  int
	}
)

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

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных")
	}
}

func AutoMigrate() {
	if err := DB.AutoMigrate(
		&Day{},
		&Attachment{},
		&Setting{},
		&User{},
	); err != nil {
		log.Fatal("Ошибка миграции таблиц")
	}
}

func RenderDatabase() {
	DB.Where("id = 1").FirstOrCreate(&Setting{
		Month: 12,
	})

	adminPass, _ := utils.HashPassword(config.Config.ADMIN_PASSWORD)
	adminRefresh, _ := utils.NewRefreshToken()

	DB.Where(&User{Email: config.Config.ADMIN_EMAIL, Role: "admin"}).FirstOrCreate(&User{
		Email:        config.Config.ADMIN_EMAIL,
		Password:     adminPass,
		RefreshToken: adminRefresh,
	})
}
