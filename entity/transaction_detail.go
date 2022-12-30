package entity

import "github.com/google/uuid"

type TransactionDetail struct {
	Id            uuid.UUID `gorm:"primaryKey;column:transaction_detail_id;type:varchar(36)"`
	SubTotalPrice int64     `gorm:"column:sub_total_price"`
	Price         int64     `gorm:"column:price"`
	Quantity      int32     `gorm:"column:quantity"`
	TransactionId uuid.UUID
	ProductId     uuid.UUID
	Product       Product
	Transaction   Transaction
}

func (TransactionDetail) TableName() string {
	return "tb_transaction_detail"
}
