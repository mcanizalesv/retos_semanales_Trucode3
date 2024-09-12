package main

import (
	"challenge5/files"
	"fmt"
)

func main() {

	// Archivo donde se guardarán los registros
	const recordsFile = "records.json" //Buscar lo que son las variables de entorno. 

	// Cargar registros desde el archivo y los guarda en el map "records"
	records := files.LoadRecords(recordsFile)

	// Mensaje de bienvenida
	fmt.Println("This is Spaced repetition app\nIt helps you to learn\nusing the spaced repetition algorithm!!!\nGood luck!")

	OptionUsers := files.PromptUserChoice()

	switch OptionUsers {
		

	case files.InvalidOption:

		fmt.Println("Thanks for using this app")

	case files.Register:
		
		NameInput:=files.PromptUserName()
		fmt.Println("The name is",NameInput)
		
		// Valida la existencia del nombre ingresado en NameInput
		NameInput =  files.NameInputValidation(NameInput,records)

		files.UsersInicialization(NameInput,records)

		// Guarda los registros actualizados en el archivo
		files.SaveRecords(records,recordsFile)

		fmt.Println("Records updated successfully!")
		fmt.Println("")
		fmt.Println("You can start playing now!")
		fmt.Println("")

		files.SpacedRepetition(NameInput,records,recordsFile)


	case files.Login:
		NameInput :=files.PromptUserName()

		// Verificar si el usuario existe en los registros
		if _, exists := records[NameInput]; exists {
			fmt.Println("User logged in successfully!")
			fmt.Println("")
		
			files.SpacedRepetition(NameInput,records,recordsFile)
		} else {
			fmt.Println("User not found.")
		}

	default:
		fmt.Println("Invalid option selected")
	}
	
}
