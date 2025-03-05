package repository

import (
	"encoding/json"
	"github.com/tidwall/buntdb"
	"log"
	"time"
)

type Data struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Repository struct {
	db *buntdb.DB
}

func NewRepository() *Repository {
	db, err := buntdb.Open("myapp.db")
	if err != nil {
		log.Fatal("Veritabanı bağlantısı kurulamadı:", err)
	}
	return &Repository{db: db}
}

func (r *Repository) Add(data Data, ttl time.Duration) error {

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return r.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(data.Id, string(jsonData), &buntdb.SetOptions{Expires: true, TTL: ttl})
		return err
	})
}

func (r *Repository) GetAll() ([]Data, error) {
	var dataList []Data

	err := r.db.View(func(tx *buntdb.Tx) error {
		return tx.Ascend("", func(key, value string) bool {
			var data Data
			if err := json.Unmarshal([]byte(value), &data); err != nil {
				return false
			}
			dataList = append(dataList, data)
			return true
		})
	})

	return dataList, err
}
