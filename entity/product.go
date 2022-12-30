package entity

import "github.com/google/uuid"

type Product struct {
	Id                 uuid.UUID           `gorm:"primaryKey;column:product_id;type:varchar(36)"`
	Name               string              `gorm:"index;column:name;type:varchar(100)"`
	Price              int64               `gorm:"column:price"`
	Quantity           int32               `gorm:"column:quantity"`
	TransactionDetails []TransactionDetail `gorm:"ForeignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (Product) TableName() string {
	return "tb_product"
}
