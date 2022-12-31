package model

type ProductModel struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    int64  `json:"price"`
	Quantity int32  `json:"quantity"`
}

type ProductCreateOrUpdateModel struct {
	Name     string `json:"name" validate:"required"`
	Price    int64  `json:"price" validate:"required"`
	Quantity int32  `json:"quantity" validate:"required"`
}
