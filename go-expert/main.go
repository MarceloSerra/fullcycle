package main

type Order struct {
	ID string
	Price float64
	Quantity int
}

func NewOrder() *Order {
	return &Order{}
}

func (o *Order) setPrice(price float64) {
	o.Price = price
}

func (o *Order) getTotal() float64{
	return o.Price * float64(o.Quantity)
}

func main(){
	order := NewOrder()
	order.ID = "01"
	order.Price = 305.50
	order.Quantity = 10

	order2 := order

	order2.setPrice(10);

	println(int(order.getTotal()))
	println(int(order2.getTotal()))
	println(int(order.getTotal()))
}