package actions

import "fmt"

type Child struct {
	Name   string
	Status Status
}

// Función que inicializa los valores del struct Person
func NewChild(name string) *Child {
	return &Child{Name: name, Status: Status{
		Hunger:     100,
		Sleepyness: 10,
		Stamina:    80},
	}
}

// Método, dependiendo del nivel de intensidad, modifica los fields del field status (Hunger,Sleepyness,Stamina) del struct child
func (c *Child) Exercising(excerciseIntensity int) string {

	switch excerciseIntensity {
	case 1:
		c.Status.Stamina -= 10
		c.Status.Hunger += 10
		return "I have exercised at low intensity."
	case 2:
		c.Status.Stamina -= 25
		c.Status.Hunger += 30
		return "I have exercised at medium intensity."
	case 3:
		c.Status.Stamina -= 50
		c.Status.Hunger += 60
		return "I have exercised at high intensity."
	default:
		return "Invalid intensity level."

	}
}

//Método aplicado al struct Child que modifica la 
func (c *Child) Sleeping() string {
	c.Status.Hunger +=  20
	c.Status.Sleepyness = 0
	c.Status.Stamina = 100
	return "I have slept."

}

func (c *Child) Studying() string {
	c.Status.Hunger += 25
	c.Status.Sleepyness += 30
	c.Status.Stamina -= 30
	return "I have studied."

}

func (c *Child) Eating(f Food) {
	c.Status.Hunger = f.Hunger(&c.Status.Hunger) * 0.8
	c.Status.Stamina = f.Stamina(&c.Status.Stamina) * 1.2
	c.Status.Sleepyness = f.Sleepiness(&c.Status.Sleepyness) * 1.1
	f.Eat()
}

func (c *Child) PrintInitChild() string {
	return fmt.Sprintf("My name is %s, this is my status:\n Hunger: %.2f\n Sleepyness: %.2f\n Stamina: %.2f\n", c.Name, c.Status.Hunger, c.Status.Sleepyness, c.Status.Stamina)
}

// Esta función imprime el estatus inicial del niño que se crea
func (c *Child) PrintHeader() {
	fmt.Print(c.PrintInitChild())
}
