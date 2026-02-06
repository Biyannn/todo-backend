package todo

import (
	"encoding/json"
	"os"
)

const filePath = "data/todos.json"

// load todos dari file
func loadFromFile() ([]Todo, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []Todo{}, nil
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var todos []Todo
	json.Unmarshal(file, &todos)
	return todos, nil
}

// simpan todos ke file
func saveToFile(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}
