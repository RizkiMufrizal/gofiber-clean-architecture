package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Id       uint8  `gorm:"primaryKey;autoIncrement:true;column:id"`
	Name     string `gorm:"index;column:name"`
	Price    int64  `gorm:"column:price"`
	Quantity int32  `gorm:"column:quantity"`
}
