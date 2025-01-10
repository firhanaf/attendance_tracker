package database

import (
	"attendance_app/apps/config"
	"attendance_app/features/attendances/data"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL(config *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&data.Attendances{})
}
