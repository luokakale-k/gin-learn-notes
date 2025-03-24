package model

type User struct {
	ID   uint   `gorm:"primaryKey" json:"id"`  // 小写字段名
	Name string `gorm:"size:50" json:"name"`   // 小写字段名
	Age  int    `gorm:"default:18" json:"age"` // 小写字段名
}
