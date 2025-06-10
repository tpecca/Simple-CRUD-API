package handlers

import (
	"CRUD-API/models"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
)

var db *gorm.DB

func InitializeDatabase(database *gorm.DB) {
	db = database
	db.AutoMigrate(&models.Item{})
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var items []models.Item
	db.Find(&items)
	json.NewEncoder(w).Encode(items)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)
	db.Create(&item)
	json.NewEncoder(w).Encode(item)
}

func PutHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)
	db.Save(&item)
	json.NewEncoder(w).Encode(item)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)
	db.Delete(&item)
	json.NewEncoder(w).Encode(item)
}
