package actions

import "fmt"

type Person struct {
	Name   string
	Status Status
}

// Función que inicializa los valores del struct Person
func NewPerson(name string) *Person {
	return &Person{Name: name, Status: Status{
		Hunger:     80,
		Sleepyness: 10,
		Stamina:    80},
	}
}

func (p *Person) Exercising(excerciseIntensity int) string {

	switch excerciseIntensity {
	case 1:
		p.Status.Stamina -= 10
		p.Status.Hunger += 10
		return "I have exercised at low intensity."
	case 2:
		p.Status.Stamina -= 25
		p.Status.Hunger += 30
		return "I have exercised at medium intensity."
	case 3:
		p.Status.Stamina -= 50
		p.Status.Hunger += 60
		return "I have exercised at high intensity."
	default:
		return "Invalid intensity level."

	}
}

func (p *Person) Sleeping() string {
	p.Status.Hunger += 20
	p.Status.Sleepyness = 0
	p.Status.Stamina = 100
	return "I have slept."

}

func (p *Person) Studying() string {
	p.Status.Hunger += 25
	p.Status.Sleepyness += 30
	p.Status.Stamina -= 30
	return "I have studied."

}

func (p *Person) Eating(f Food) {
	p.Status.Hunger = f.Hunger(&p.Status.Hunger)
	p.Status.Stamina = f.Stamina(&p.Status.Stamina)
	p.Status.Sleepyness = f.Sleepiness(&p.Status.Sleepyness)
	f.Eat()
}

func (p *Person) PrintStatus() string {
	return fmt.Sprintf("The actual status for %s is:\n Hunger:%.2f\n Sleep:%.2f\n Stamina:%.2f\n", p.Name, p.Status.Hunger, p.Status.Sleepyness, p.Status.Stamina)
}

func (p *Person) PrintInitPerson() string {

	return fmt.Sprintf("The actual status for %s is:\n Hunger:%.2f\n Sleep:%.2f\n Stamina:%.2f\n", p.Name, p.Status.Hunger, p.Status.Sleepyness, p.Status.Stamina)

}

//Esta función imprime el estatus inicial de la persona  que se crea
func (p *Person) PrintHeader() {
	fmt.Print(p.PrintInitPerson())
}
