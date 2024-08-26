package main

import (
	"fmt"

	"challenge3/actions"
)

type person_child int

const (
	Unselected = iota
	Person
	Child
)

type personAction int

const (
	Action_Unselected personAction = iota
	Exercise
	Sleep
	Study
	Eat
	Show_status
)

type excerciseIntensity int

const (
	InvalidIntensity excerciseIntensity = iota
	Low
	Medium
	Hard
)

type foodOptions int

const (
	Unselected_Food foodOptions = iota
	Dessert 
	Meal
	Energizer
)

func main() {

	var action actions.Actions
	var personAction personAction
	var personChild person_child
	var childName string
	var personName string
	var intensity excerciseIntensity

	fmt.Println("What did you want to create? Enter 1 for a person or 2 for a child")
	_, err := fmt.Scan(&personChild)
	if err != nil {
		personChild = Unselected
	}

	switch personChild {

	case Person:
		fmt.Println("What is the person name?")
		_, err := fmt.Scan(&personName)
		if err != nil {
			fmt.Println("Write a name")
		}
		personStruct := actions.NewPerson(personName)
		personStruct.PrintHeader()
		action = personStruct

	case Child:
		fmt.Println("What is the child name?")
		_, err := fmt.Scan(&childName)
		if err != nil {
			fmt.Println("Write a name")
		}
		childStruct := actions.NewChild(childName)
		childStruct.PrintHeader()
		action = childStruct

	default:
		personChild = Unselected

	}

	for {
		fmt.Println("Choose an action for your person or child")
		fmt.Println("Choose:\n 0.Exit\n 1.Exercise\n 2.Sleep\n 3.Study\n 4.Eat\n 5.Show Status\n ")

		_, err1 := fmt.Scan(&personAction)
		if err1 != nil {
			personAction = Unselected
		}

		switch personAction {
		case Unselected:
			fmt.Println("Thanks for use this app")

		case Exercise:
			fmt.Println("Enter the exercise intensity:\n 1.Low\n 2.Medium\n 3.Hard")
			fmt.Scan(&intensity)
			if action != nil {
				action := action.Exercising(int(intensity))
				fmt.Println(action)
			} else {
				fmt.Println("Action not initialized")
			}
		case Sleep:

			if action != nil {
				action := action.Sleeping()
				fmt.Println(action)
			} else {
				fmt.Println("Action not initialized")
			}

		case Study:

			if action != nil {
				action := action.Studying()
				fmt.Println(action)
			} else {
				fmt.Println("Action not initialized")
			}
		case Eat:
			fmt.Println("Choose the food you want to eat:\n 1.Dessert\n 2.Meal\n 3.Energizer")
			var foodOption foodOptions
			_, err := fmt.Scan(&foodOption)
			if err != nil {
				fmt.Println("Invalid input for food option")
			}
			switch foodOption {
			case Dessert:
				dessert := actions.Dessert{Name: "Ice Cream", Taste: "Delicious"}
				action.Eating(dessert)
			case Meal:
				meal := actions.Meal{Name: "Pizza", Taste: "Delicious"}
				action.Eating(meal)
			case Energizer:
				energizer := actions.Energizer{Name: "Red Bull", Taste: "Delicious"}
				action.Eating(energizer)
			default:
				fmt.Println("This food doesn't exists")
			}
		case Show_status:
			action.PrintHeader()
		default:
			fmt.Println("This action doesn't exists")
			//falta agregar el default
		}

		if personAction == Action_Unselected {
			break
		}
	}

}
