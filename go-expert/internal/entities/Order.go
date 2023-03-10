package entities

type Order struct {
	ID string
	Price float64
	Quantity int
}

func NewOrder() *Order {
	return &Order{}
}

func (o Order) getTotal() float64{
	return o.Price * float64(o.Quantity)
}