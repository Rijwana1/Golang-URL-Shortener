package db

import (
	"poc/url-shortener/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateConnection() *gorm.DB {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(
		"postgres",
		"host=localhost user=postgres port=5432 dbname=postgres password=rijwana sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.URLTable{})
	return db
}
