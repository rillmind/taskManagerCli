package task

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/rillmind/taskManagerCli/enum/status"
)

type Task struct {
	ID          int           `json:"id"`
	Description string        `json:"description"`
	Status      status.Status `json:"status"`
	CreatedAt   string        `json:"createdAt"`
	UpdatedAt   string        `json:"updatedAt"`
}

func CreateTask(description string) Task {
	var tasks []Task
	id := 0

	_, err := os.Stat("tasks.json")
	if !os.IsNotExist(err) {
		tasks = readJsonFile("tasks.json")
		if len(tasks) > 0 {
			id = tasks[len(tasks)-1].ID
		}
	}

	task := Task{
		ID:          id + 1,
		Description: description,
		Status:      status.TODO,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	tasks = append(tasks, task)
	createJsonFile(tasks)

	return task
}

func createJsonFile(tasks any) {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Erro ao gerar JSON:", err)
	}

	err = os.WriteFile("tasks.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
	}

	fmt.Println("Arquivo pessoas.json criado com sucesso!")
}

func readJsonFile(filename string) []Task {
	jsonFile, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	var tasks []Task
	err = json.Unmarshal(jsonFile, &tasks)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	return tasks
}
