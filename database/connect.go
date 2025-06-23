package database

import (
	"api/config"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

type Database struct {
	DB *gorm.DB
}

func NewDatabaseConnection() (*Database, error) {
	var err error

	dsn := "host=" + config.DB_HOST + " user=" + config.DB_USER + " password=" + config.DB_PASSWORD + " dbname=" + config.DB_NAME + " port=5432 sslmode=disable TimeZone=Asia/Manila"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected and migrated successfully")

	return &Database{DB: db}, nil
}

func (d *Database) Close() error {
	sqlDB, err := d.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

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
