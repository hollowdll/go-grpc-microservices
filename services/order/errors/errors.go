package errors

import (
	"errors"
)

var (
	ErrGetProductPrices             = errors.New("cannot get prices of the ordered products")
	ErrCheckProductStockQuantities  = errors.New("cannot check the stock quantities of the ordered products")
	ErrNotEnoughProducts            = errors.New("not enough ordered products in stock")
	ErrCreatePayment                = errors.New("failed to process payment")
	ErrReduceProductStockQuantities = errors.New("cannot reduce stock quantities of the ordered products")
)
