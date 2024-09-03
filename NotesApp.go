package main

import (
	"bufio"
	"com.github.rakhmedovrs/udemy-go-course/note"
	"fmt"
	"os"
	"strings"
)

func notesApp() {
	userNote := *getNote()
	userNote.Display()
	err := userNote.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}
	fmt.Println("Note saved successfully")
}

func getNote() *note.Note {
	title := getUserInputs("Note title: ")
	content := getUserInputs("Note content: ")
	userNote, err := note.New(title, content)
	if err != nil {
		panic(err)
	}
	return userNote
}

func getUserInputs(prompt string) string {
	fmt.Println(prompt)
	value, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return ""
	}
	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")
	return value
}
