package store

import (
	"fmt"
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
	out := &model.URLTable{}
	fmt.Println(shortURL)
	err := db.Raw(qryFind, shortURL).Scan(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

var qryFind = `
SELECT
	short_url,
 	long_url 
FROM
	 url_tables 
WHERE
	short_url = ?
`
