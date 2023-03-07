package main

import "fmt"

//THIS IS A STRUCT
type Order struct {
	ID       string
	Price    float64
	Quantity int
}

func (o Order) SetPrice(price float64) {
	o.Price = price
	fmt.Println("Price dentro do SetPrice: ", o.Price)
}

//THIS IS A METHOD FOR STRUCT
func (o Order) geTotal() float64 {
	return o.Price * float64(o.Quantity)
}

//THIS RETURN THIS VALUE CALCULATE IN THE METHOD
func main() {
	order := Order{
		ID:       "123",
		Price:    10.0,
		Quantity: 5,
	}
	order.SetPrice(20.0)
	fmt.Println("Preço total: ", order.geTotal())
}
