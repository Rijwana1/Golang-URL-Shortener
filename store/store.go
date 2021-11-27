package store

import (
	"poc/url-shortener/db"
	"poc/url-shortener/model"
)

func Create(url model.URLTable) error {
	db := db.CreateConnection()
	defer db.Close()
	err := db.Create(&url).Error
	if err != nil {
		return err
	}
	return nil
}

func Find(shortURL string) (*model.URLTable, error) {
	db := db.CreateConnection()
	defer db.Close()
	url := &model.URLTable{}
	err := db.First(&url, shortURL).Scan(&url).Error
	if err != nil {
		return nil, err
	}
	return url, nil
}
