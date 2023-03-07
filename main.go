package main

import "fmt"

//THIS IS A STRUCT
type Order struct {
	ID       string
	Price    float64
	Quantity int
}

// '*' aponta para o valor na memória
func (o *Order) SetPrice(price float64) {
	o.Price = price
	fmt.Println("Price dentro do SetPrice: ", o.Price)
}

//THIS IS A METHOD FOR STRUCT
func (o Order) geTotal() float64 {
	return o.Price * float64(o.Quantity)
}

// '*' e '&' estão recebendo um ponteiro, ou seja, order2 e order estão apontando para o mesmo local na memoria
func NewOrder() *Order {
	return &Order{}
}

//THIS RETURN THIS VALUE CALCULATE IN THE METHOD
func main() {
	order := NewOrder()
	order.Quantity = 10
	order.ID = "123"
	order.Price = 5.0
	order.SetPrice(20.0)

	order2 := order // está apontando para onde está guardado
	order.Price = 5
	fmt.Println("Preço total: ", order.geTotal())
	fmt.Println(order2.Price)
}
