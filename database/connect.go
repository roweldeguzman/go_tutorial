package database

import (
	"api/config"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Open() error {

	var err error

	//dsn := config.DB_USER + ":" + config.DB_PASSWORD + "@tcp(" + config.DB_HOST + ":" + config.DB_PORT + ")/" + config.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "host=" + config.DB_HOST + " user=" + config.DB_USER + " password=" + config.DB_PASSWORD + " dbname=" + config.DB_NAME + " port=5432 sslmode=disable TimeZone=Asia/Manila"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
		},
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	DB, _ := DB.DB()
	return DB.Close()
}
