[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-22041afd0340ce965d47ae6ef1cefeee28c7c493a6346c4f15d667ab976d596c.svg)](https://classroom.github.com/a/kyyW2ZP_)
# Weekend challenge 2

You will code a program to run an automated coffee store.

In this store, the customer creates their drink by choosing different options through a series of steps. The customer can cancel any of the steps to go back and choose a different option.
When the drink is complete, the customer enters their order and it will be delivered in a FIFO order.
The customer's name is a unique identifier so if the customer places an order again, it is invalidated, indicating that they cannot enter a new order until the pending one is delivered.

The main function should be able to cover succesfully the next set of instructions:

```go
// Create a new orderQueue
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
fmt.Printf("%s placed an order, the required drink has the following ingredients: %v\n", order1.Name, order1.Drink.ListIngredients())

// Shows error because a client1 order already exists
fmt.Printf("%s repeats the order, the required drink has the following ingredients: %v\n", order1.Name, order1.Drink.ListIngredients())
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
fmt.Printf("%s placed an order, the required drink has the following ingredients: %s\n", order2.Name, order2.Drink.ListIngredients())

// Deliver the coffees in orderQueue
nextOrder,_ := orderManager.Next()
fmt.Printf("Coffee for %s ready to be tasted!!!\n", nextOrder.Name)
nextOrder,_ = orderManager.Next()
fmt.Printf("Coffee for %s ready to be tasted!!!\n", nextOrder.Name)
// Show an error message if there are no more orders to be served
if _, err := orderManager.Next(); err != nil {
    fmt.Println(err)
}
```

The program can be tested either with just a `go run` or ideally with an executable. So we do something like:

```shell
go run .
```

or

```shell
./w2challenge
```

## Recommended project structure

```md
project-root-directory/
├── go.mod
├── main.go
├── data_structures/
│ ├── set.go
│ ├── stack.go
├── models/
│ ├── client.go
│ ├── drink.go
│ ├── order_manager.go
│ ├── order.go
```

## Recommended structures

### data_structures/set

#### Fields

| Name  | Type            |
| ----- | --------------- |
| ìtems | map[string]bool |

#### Methods

**Append** - Receives an string, and stores it on a map key. If the item exists, returns the "item already exists" error, else stores the item and returns nil.

**Exists** - Receives an string, returns the existence of the item in a map.

**Delete** - Receives an string, deletes the item which key matches with the received string.

### data_structures/stack

#### Fields

| Name  | Type     |
| ----- | -------- |
| ìtems | []string |

#### Methods

**Push** - Receives an string, and appends it to its items collection.

**Pop** - Removes the last item of its items collection and returns it, returns the error "no more items in the stack" if the collection's length is 0.

**ToString** - Transforms its items collection to a string look at `strings.Join`.

### models/client

#### Fields

| Name | Type   |
| ---- | ------ |
| Name | string |

### models/drink

#### Fields

| Name        | Type  |
| ----------- | ----- |
| ingredients | Stack |

#### Methods

**AddIngredient** - Receives an string, and pushes it to its ingredients stack.

**RemoveIngredient** - Pops from its ingredients stack if it returns an error the method returns "no more ingredients in the list", otherwise returns the popped item.

**ListIngredients** - Returns the call to the method ToString of its ingredients stack.

### models/order_manager

#### Fields

| Name    | Type    |
| ------- | ------- |
| orders  | []Order |
| clients | Set     |

#### Methods

**enqueue** - Receives an Order, and appends it to its orders slice.

**dequeue** - If the orders length is 0, returns the error "no more orders in the stack", otherwise returns the first item in the orders.

**Add** - Receives an order, if the order's client name already exists in the clients set returns the error "{client name}'s order already exists". Otherwise appends the client's name to the clients set and enqueues the order.

**Next** - Calls a dequeue, if it returns an error, this method should return "no more orders to be served", otherwise deletes the client's name from the clients set and returns the dequeued order.

### models/order

#### Fields

| Name   | Type   |
| ------ | ------ |
| Client | Client |
| Drink  | Drink  |

\*In this structure recall to the anonymous fields topic.
