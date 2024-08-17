package models

import (
	datastructures "trucode.app/w2challenge/data_structures"
)

type Drink struct {
	ingredients datastructures.Stack
}

func (in *Drink) AddIngredient(value string) {
	in.ingredients.Push(value)
}


func (in *Drink) RemoveIngredient() (string, error) {
	return in.ingredients.Pop()

}

func (in *Drink) ListIngredients(separator string) string {
	return in.ingredients.ToString(separator)
}