package order

import "github.com/nhtuan0700/orders-api/model"

type FindAllPage struct {
	Size   uint
	Offset uint
}

type FindResult struct {
	Orders []model.Order
	Cursor uint64
}
