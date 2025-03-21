package model

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:50"`
	Age  int    `gorm:"default:18"`
}
