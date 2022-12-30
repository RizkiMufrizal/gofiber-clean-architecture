package entity

import "github.com/google/uuid"

type Transaction struct {
	Id                 uuid.UUID           `gorm:"primaryKey;column:transaction_id;type:varchar(36)"`
	TotalPrice         int64               `gorm:"column:total_price"`
	TransactionDetails []TransactionDetail `gorm:"ForeignKey:TransactionId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (Transaction) TableName() string {
	return "tb_transaction"
}
