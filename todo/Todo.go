package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("text must be provided")
	}

	return &Todo{
		Text: text,
	}, nil
}

func (n Todo) Display() {
	fmt.Println(n.Text)
}

func (n Todo) Save() error {
	fileName := "todo.json"
	bytes, err := json.Marshal(n)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, bytes, 0644)
}
