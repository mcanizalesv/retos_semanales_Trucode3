package main

import (
	"fmt"

	datastructures "trucode.app/w2challenge/data_structures"
	models "trucode.app/w2challenge/models"
)

func main(){
	/* This set of instructions shows how a Set works -and the code below how you can implement it-.
	 * You will need to implement it in order to accomplish the project requirements. 
	 * Make sure to correctly isolate the code and that it handles errors in a good manner.
	 */ 
	set := datastructures.Set{}
	// {}
	set.Append("1")
	// {"1":true}
	set.Append("2")
	// {"1":true, "2": true}
	set.Append("3")
	// {"1":true, "2":true, "3":true}
	set.Append("1")
	// {"1":true, "2":true, "3":true}
	set.Append("1")
	// {"1":true, "2":true, "3":true}
	set.Append("1")
	// {"1":true, "2":true, "3":true}
	fmt.Println(set)
	// map{1:true 2:true 3:true}
	fmt.Println(set.Exists("2"))
	// true
	set.Delete("2")
	// {"1":true, "3":true}
	fmt.Println(set.Exists("2"))
	// false
	fmt.Println(set)
	// map{1:true 3:true}
	

	var orderManager models.OrderManager

	// Create a client
	client1 := models.Client{Name: "John"}
	fmt.Printf("%s wants to order\n", client1.Name)
	// Create a drink for client1
	var drink1 models.Drink

	// Add ingredients
	drink1.AddIngredient("coffee")
	drink1.AddIngredient("water")
	drink1.AddIngredient("milk")
	// Remove the last ingredient
	drink1.RemoveIngredient()
	// Keep adding ingredients
	drink1.AddIngredient("lactose-free milk")
	drink1.AddIngredient("syrup")

	// Append a new order to the orderQueue
	order1 := models.Order{Client: client1, Drink: drink1}
	orderManager.Add(order1)
	fmt.Printf("%s placed an order, the required drink has the following ingredients: %v\n", order1.Client.Name, order1.Drink.ListIngredients(" "))

	// Shows error because a client1 order already exists
	fmt.Printf("%s repeats the order, the required drink has the following ingredients: %v\n", order1.Client.Name, order1.Drink.ListIngredients(" "))
	if err := orderManager.Add(models.Order{Client: client1, Drink: drink1}); err != nil {
		fmt.Println(err)
	}

	// Create another client
	client2 := models.Client{Name: "Jane"}

	// Create a drink for client2
	var drink2 models.Drink

	// Add ingredients
	drink2.AddIngredient("coffee")
	drink2.AddIngredient("water")
	drink2.AddIngredient("milk")
	// Remove the last ingredient
	drink2.RemoveIngredient()

	// Append a new order to the orderQueue
	order2 := models.Order{Client: client2, Drink: drink2}
	orderManager.Add(order2)
	fmt.Printf("%s placed an order, the required drink has the following ingredients: %s\n", order2.Client.Name, order2.Drink.ListIngredients(" "))

	//Deliver the coffees in orderQueue
	nextOrder,_ := orderManager.Next()
	fmt.Printf("Coffee for %s ready to be tasted!!!\n", nextOrder.Client.Name)
	nextOrder,_ = orderManager.Next()
	fmt.Printf("Coffee for %s ready to be tasted!!!\n", nextOrder.Client.Name)
	// Show an error message if there are no more orders to be served
	if _, err := orderManager.Next(); err != nil {
		fmt.Println(err)
	}
}

/* A set is a data structure that stores unique elements
 * of the same type in a sorted order. Each value is a key,
 * which means that we access each value using the value itself.
 * Accordingly, each value in a set must be unique.
 */
// type Set struct{
// 	// A common way to configure a set is to have a key with desired store value and a boolean indicating it exists
// 	items map[?]?
// }

// func (s *Set) append(value string){
// 	// What's the current state of items?
// 	// If it doesn't exist, how can you initialize it?
// 	s.items[value] = true
// }

// func (s Set) exists(value string) bool {
// 	// We can get a key from the map, remember the return come as (value, exists)
// 	s.items[value]
// }

// func (s *Set) delete(value string) {
// 	// Remember there is a method called delete, how can you use it here?
// }

