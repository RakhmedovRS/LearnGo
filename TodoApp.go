package main

import (
	"bufio"
	"com.github.rakhmedovrs/udemy-go-course/todo"
	"fmt"
	"os"
	"strings"
)

func todoApp() {
	userTodo := *getTodo()
	err := outputData(userTodo)
	if err != nil {
		return
	}
}

func getTodo() *todo.Todo {
	content := getUserTodoInputs("Todo content: ")
	userTodo, err := todo.New(content)
	if err != nil {
		panic(err)
	}
	return userTodo
}

func getUserTodoInputs(prompt string) string {
	fmt.Println(prompt)
	value, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return ""
	}
	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")
	return value
}
