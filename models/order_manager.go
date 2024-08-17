package models

import (
	"fmt"

	datastructures "trucode.app/w2challenge/data_structures"
)

type OrderManager struct {
	orders []Order
	clients datastructures.Set
}

func (ord *OrderManager) Enqueue(value Order) {
	ord.orders = append(ord.orders, value)
}

func (ord *OrderManager) Dequeue() (Order, error) {
	if len(ord.orders) == 0 {
		return Order{} ,fmt.Errorf("no more orders in the queue")
	}
	value := ord.orders[0]
	ord.orders = ord.orders[1:]
	return value, nil
}

func (ord *OrderManager) Add(order Order) error {
	if ord.clients.Exists(order.Client.Name) {
		return fmt.Errorf("%s's order already exists", order.Client.Name)
	}

	ord.clients.Append(order.Client.Name)

	ord.orders = append(ord.orders, order)
	return nil
}

func (ord *OrderManager) Next() (Order,error) {
	order, err := ord.Dequeue()
	if err != nil {
		return Order{}, fmt.Errorf("no more orders to be served")
	}
	ord.clients.Delete(order.Client.Name)
	dequededOrder := ord.orders[0]
    ord.orders = ord.orders[1:]
	return dequededOrder,nil
	
	
}