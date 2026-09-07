package out

import "github.com/shopspring/decimal"

type ItemAvailability struct {
	ProductID int64
	Name      string
	Price     decimal.Decimal
	Quantity  int32
	Available int
}
