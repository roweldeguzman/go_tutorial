package database

import (
	"api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Connect *gorm.DB

func Open() error {

	var err error

	dsn := config.DB_USER + ":" + config.DB_PASSWORD + "@tcp(" + config.DB_HOST + ":" + config.DB_PORT + ")/" + config.DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	Connect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	DB, _ := Connect.DB()
	return DB.Close()
}
