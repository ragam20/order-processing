package out

type ItemAvailability struct {
	ProductID   int64
	Name        string
	Price       int64
	Quantity    int32
	Available   int
	IsAvailable bool
}
