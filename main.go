package main

import "fmt"

var applications = map[int]func(){
	1: investmentCalculator,
	2: bank,
	3: notesApp,
	4: todoApp,
}

type saver interface {
	Save() error
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving:", err)
		if err != nil {
			return err
		}
	}
	fmt.Println("Saved successfully")
	return nil
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func main() {
	fmt.Println("Which application do you want to run?")
	fmt.Println("1. Investment calculator")
	fmt.Println("2. Bank")
	fmt.Println("3. Notes app")
	fmt.Println("4. Todo app")
	var userInput int
	fmt.Scanf("%d", &userInput)
	application, ok := applications[userInput]
	if !ok {
		fmt.Printf("Your choice %d is not recognized", userInput)
		return
	}

	application()
}
