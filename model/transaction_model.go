package model

import "github.com/google/uuid"

type TransactionModel struct {
	Id                 string                   `json:"id"`
	TotalPrice         int64                    `json:"total_price"`
	TransactionDetails []TransactionDetailModel `json:"transaction_details"`
}

type TransactionDetailModel struct {
	Id            string    `json:"id"`
	SubTotalPrice int64     `json:"sub_total_price"`
	Price         int64     `json:"price"`
	Quantity      int32     `json:"quantity"`
	TransactionId uuid.UUID `json:"transaction_id"`
	ProductId     uuid.UUID `json:"product_id"`
	Product       ProductModel
}
