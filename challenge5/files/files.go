package files

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

type Progress struct {
	Apparitions     float64 `json:"apparitions"`
	Correct_answers float64 `json:"correct_answers"`
	Retention_rate  float64 `json:"retention_rate"`
}

type Words struct {
	Learning    Progress `json:"learning"`
	Walking     Progress `json:"walking"`
	Certain     Progress `json:"certain"`
	Intricate   Progress `json:"intricate"`
	Calculation Progress `json:"calculation"`
}

type User struct {
	Words Words `json:"words"`
}

type optionUsers int

const (
	InvalidOption optionUsers = iota
	Register 
	Login
)

//PromptUserName() recibe el nombre del usuario a través de un fmt.Scan y lo asigna a la variable NameInput
func PromptUserName() string {
	var NameInput string

	fmt.Println("Please tell me your name?")
	_, err := fmt.Scan(&NameInput) 
	if err != nil {
		fmt.Println("Error reading input:", err)
		
	}
	NameInput = strings.ReplaceAll(NameInput, " ", "")
	NameInput = strings.ToLower(NameInput)
	
	return NameInput
}
// PrompUserChoice() permite realizar la selección de la opción por parte del usuario ("register","login" o "exit")
func PromptUserChoice() optionUsers {
	var OptionUsers optionUsers

	fmt.Println("Did you want to (into the num, for example, 1 or 2):\n0.Exit\n1.Register\n2.Login")
	_, err := fmt.Scan(&OptionUsers)

	if err != nil {
		OptionUsers = InvalidOption
	}

	return OptionUsers
}

// NameInputValidation busca entre los registros del map que guarda los datos traidos del JSON y verifica que no exista una clave igual a la entrada en NameInput.
func NameInputValidation (NameInput string,records map[string]User) string {
	for {
		if _, exists := records[NameInput]; exists {
			fmt.Printf("This name %s already exists. Please, enter a different name: ", NameInput)
			_, err := fmt.Scan(&NameInput)
			if err != nil {
				fmt.Println("Incorrect input. Please try again.")
				continue 
			}
		} else {
			break 
		}
	}
	return NameInput 

}

// UsersInicialization() inicializa un nuevo usuario en el map records
func UsersInicialization(NameInput string, records map[string]User)  {
	users := User{Words: Words{
		Learning:    Progress{},
		Walking:     Progress{},
		Certain:     Progress{},
		Intricate:   Progress{},
		Calculation: Progress{},
	}}

	// Guarda el nuevo registro en el mapa
	records[NameInput] = users
}

// LoadRecords sirve para cargar los registros desde un archivo
 func LoadRecords(recordsFile string) map[string]User {
	file, err := os.Open(recordsFile)
	if err != nil {
		if os.IsNotExist(err) {
			// Si el archivo no existe, devolvemos un mapa vacío - crear archivo
			return make(map[string]User)
		}

		log.Fatal(err)
	}

	defer file.Close()

	var recordsMap map[string]User

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	//Revisar con return early - refactorización
	if len(fileBytes) == 0{
		return make(map[string]User)
	}

    if err := json.Unmarshal(fileBytes, &recordsMap); err != nil {log.Fatal(err)}
	
	return recordsMap
}


// SaveRecords para guardar los registros en un archivo
func SaveRecords(records map[string]User, recordsFile string) {
	bytes, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		log.Fatal(fmt.Errorf("marshall records:%w",err)) //Evaluar si está bien que falle toda la app si se hace mal la lectura
	}
	err = os.WriteFile(recordsFile, bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

//UpdateProgress() actaliza los valores del Progress de cada Word
func UpdateProgress(words *Words, word string, correct bool) {
	switch word {
	case "learning":
		words.Learning.Apparitions++
		if correct {
			words.Learning.Correct_answers++
		}
		words.Learning.Retention_rate = words.Learning.Correct_answers / words.Learning.Apparitions
	case "walking":
		words.Walking.Apparitions++
		if correct {
			words.Walking.Correct_answers++
		}
		words.Walking.Retention_rate = words.Walking.Correct_answers / words.Walking.Apparitions
	case "certain":
		words.Certain.Apparitions++
		if correct {
			words.Certain.Correct_answers++
		}
		words.Certain.Retention_rate = words.Certain.Correct_answers / words.Certain.Apparitions
	case "intricate":
		words.Intricate.Apparitions++
		if correct {
			words.Intricate.Correct_answers++
		}
		words.Intricate.Retention_rate = words.Intricate.Correct_answers / words.Intricate.Apparitions
	case "calculation":
		words.Calculation.Apparitions++
		if correct {
			words.Calculation.Correct_answers++
		}
		words.Calculation.Retention_rate = words.Calculation.Correct_answers / words.Calculation.Apparitions
	}
}

//PrintProgress imprime el retention rate de las palabras asignadas al usuario loggeado
func PrintProgress(words Words) {
    fmt.Println("Your current progress is:")
    fmt.Println("Words       | Retention Rate")

    // Formatea cada palabra y su tasa de retención
    fmt.Printf("%s | %.0f%%\n", "learning", words.Learning.Retention_rate * 100)
    fmt.Printf("%s | %.0f%%\n", "walking", words.Walking.Retention_rate * 100)
    fmt.Printf("%s | %.0f%%\n", "certain", words.Certain.Retention_rate * 100)
    fmt.Printf("%s | %.0f%%\n", "intricate", words.Intricate.Retention_rate * 100)
    fmt.Printf("%s | %.0f%%\n", "calculation", words.Calculation.Retention_rate * 100)
}


//SpacedRepetition() guarda la lógica del juego
func SpacedRepetition(NameInput string,records map[string]User,recordsFile string) {
	selectedMenuLearn := 1

	round := 0					
	wordsEnglish := []string{"learning", "walking", "certain", "intricate", "calculation"}
	englishSpanish := map[string]string{"learning": "aprender", "walking": "caminar", "certain": "cierto", "intricate": "intricado", "calculation": "calculo"}


	correctEnglishSpanish := map[string]bool{"learning": false, "walking": false, "certain": false, "intricate": false, "calculation": false}

	var greatBadWordInput string
	var WordApparition int
	InfoUser := records[NameInput]
	
	
	for selectedMenuLearn == 1 {
		fmt.Println(greatBadWordInput + " What did you want to do?")
		fmt.Println("1. Go to the next word")	
		fmt.Println("2. See my progress")	
		fmt.Println("3. Finish")	

		round++ // para ver nro de round
		fmt.Scanln(&selectedMenuLearn)
		fmt.Println("")
		if selectedMenuLearn == 1 {
			fmt.Println("round:", round ," selectedMenuLearn:", selectedMenuLearn)
		}
						
		switch selectedMenuLearn {			
			case 1:
				i := 0				
				var wordInput string	
				wordAsk := wordsEnglish[i]
				fmt.Printf("What does %s means in Spanish?\n", wordAsk)
				fmt.Scanln(&wordInput)	
				wordSpanish := englishSpanish[wordsEnglish[i]] // aprender
				
				if wordInput != wordSpanish {
					correctEnglishSpanish[wordAsk] = false
					temp := wordAsk
					wordsEnglish[i] = wordsEnglish[i+1]   
					wordsEnglish[i+1] = temp
					greatBadWordInput = "Bad!"
				} else {

					//Actualizamos campo de la palabra que recién aparece
					UpdateProgress(&InfoUser.Words, wordAsk, true)
					correctEnglishSpanish[wordAsk] = true	

					// Switch para actualizar solo la palabra actual
					switch wordAsk {
					case "learning":
						WordApparition = int(InfoUser.Words.Learning.Apparitions)
					case "walking":
						WordApparition = int(InfoUser.Words.Walking.Apparitions)
					case "certain":
						WordApparition = int(InfoUser.Words.Certain.Apparitions)
					case "intricate":
						WordApparition = int(InfoUser.Words.Intricate.Apparitions)
					case "calculation":
						WordApparition = int(InfoUser.Words.Calculation.Apparitions)
					}

					//Se actualizan las posiciones del array wordsEnglish
					newPosition := int(math.Pow(2, float64(WordApparition)))
					wordsEnglishcOPY := make([]string, len(wordsEnglish))
					copy(wordsEnglishcOPY, wordsEnglish)
		
					if len(wordsEnglish) > newPosition {
						wordsEnglish[newPosition] = wordsEnglish[0]
						for i := 0; i < newPosition; i++ {
							wordsEnglish[i] = wordsEnglishcOPY[i+1]
						}	
					} else {
						wordsEnglish[len(wordsEnglish)-1] = wordsEnglish[0]
						for i := 0; i < len(wordsEnglish) - 1; i++ {
							wordsEnglish[i] = wordsEnglishcOPY[i+1]
						}	
					}			
					greatBadWordInput = "Great!"	
					// Aquí se guardan los registros actualizados en el json
					records[NameInput] = InfoUser
					SaveRecords(records, recordsFile)	
				}	
				fmt.Println("")	
				
				case 2:	
				//muestra el progreso del usuario y vuelve a mostrar el menu
				PrintProgress(InfoUser.Words)
				selectedMenuLearn = 1
			default:		
			fmt.Println("Thank you for use this app!")
					
		}			


	}
	
	}