package models

type Item struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255"`
	Price float64
}
