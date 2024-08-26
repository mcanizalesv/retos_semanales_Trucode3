package actions

import "fmt"

// Food represents various types of food
type Food interface {
	Eat()
	Effects
}

// Meal represents a meal food
type Meal struct {
	Name  string
	Taste string
}

func (Meal) Hunger(h *float64) float64 {
	return *h - 50
}
func (Meal) Stamina(s *float64) float64 {
	return *s + 25
}
func (Meal) Sleepiness(sl *float64) float64 {
	return *sl + 10
}

// Eat implements the Eat method for Meal
func (m Meal) Eat() {

	fmt.Printf("That %s tasted %s!\n", m.Name, m.Taste)
	fmt.Println("I have eaten.")
}

// Energizer represents an energizer food
type Energizer struct {
	Name  string
	Taste string
}

func (Energizer) Hunger(h *float64) float64 {
	return *h - 10
}
func (Energizer) Stamina(s *float64) float64 {
	return *s + 50
}

func (Energizer) Sleepiness(sl *float64) float64 {
	return *sl - 50
}

// Eat implements the Eat method for Energizer
func (e Energizer) Eat() {
	fmt.Printf("That %s tasted %s!\n", e.Name, e.Taste)
	fmt.Println("I have eaten.")
}

type Dessert struct {
	Name  string
	Taste string
}

func (Dessert) Hunger(h *float64) float64 {
	return *h - 20
}
func (Dessert) Stamina(s *float64) float64 {
	return *s + 10
}
func (Dessert) Sleepiness(sl *float64) float64 {
	return *sl
}

func (d Dessert) Eat() {
	fmt.Printf("That %s tasted %s!\n", d.Name, d.Taste)
	fmt.Println("I have eaten.")
}

type Effects interface {
	Hunger(*float64) float64
	Stamina(*float64) float64
	Sleepiness(*float64) float64
}
