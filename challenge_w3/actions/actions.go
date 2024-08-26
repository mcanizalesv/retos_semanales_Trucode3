package actions

type Actions interface {
	Exercising(int) string
	Sleeping() string
	Studying() string
	PrintHeader()
	Eating(Food)
}
