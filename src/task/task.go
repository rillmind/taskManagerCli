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

	_, err := os.Stat("./tasks.json")
	if !os.IsNotExist(err) {
		tasks = readJsonFile("./tasks.json")
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

func ListTasks() []Task {
	var tasks []Task

	_, err := os.Stat("./tasks.json")
	if !os.IsNotExist(err) {
		tasks = readJsonFile("./tasks.json")

		data, err := json.MarshalIndent(tasks, "", "    ")

		if err != nil {
			fmt.Println("Erro formating JSON:", err)
			return nil
		}

		fmt.Println(string(data))
	}

	return tasks
}

func ListTasksByStatus(status status.Status) []Task {
	var tasks []Task
	var statusTasks []Task

	_, err := os.Stat("./tasks.json")
	if !os.IsNotExist(err) {
		tasks = readJsonFile("./tasks.json")

		for _, task := range tasks {
			if task.Status == status {
				statusTasks = append(statusTasks, task)
			}
		}

		data, err := json.MarshalIndent(statusTasks, "", "    ")

		if err != nil {
			fmt.Println("Erro formating JSON:", err)
			return nil
		}

		fmt.Println(string(data))
	}

	return statusTasks
}

func UpdateTask(id int, description string) {
	var newTasks []Task

	tasks := readJsonFile("./tasks.json")

	for _, task := range tasks {
		if id != task.ID {
			newTasks = append(newTasks, task)
		} else if id == task.ID {
			newTask := Task{
				ID:          task.ID,
				Description: description,
				Status:      task.Status,
				CreatedAt:   task.CreatedAt,
				UpdatedAt:   time.Now().String(),
			}
			newTasks = append(newTasks, newTask)
		}
	}

	createJsonFile(newTasks)
}

func DeleteTask(id int) {
	var newTasks []Task

	tasks := readJsonFile("./tasks.json")

	for _, task := range tasks {
		if id != task.ID {
			newTasks = append(newTasks, task)
		}
	}

	createJsonFile(newTasks)
}

func Mark(id int, status status.Status) {
	var newTasks []Task

	tasks := readJsonFile("./tasks.json")

	for _, task := range tasks {
		if id != task.ID {
			newTasks = append(newTasks, task)
		} else if id == task.ID {
			newTask := Task{
				ID:          task.ID,
				Description: task.Description,
				Status:      status,
				CreatedAt:   task.CreatedAt,
				UpdatedAt:   time.Now().String(),
			}
			newTasks = append(newTasks, newTask)
		}
	}

	createJsonFile(newTasks)
}

func createJsonFile(tasks any) {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")

	if err != nil {
		fmt.Println("Error generating JSON:", err)
	}

	err = os.WriteFile("./tasks.json", jsonData, 0777)

	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func readJsonFile(filename string) []Task {
	var tasks []Task

	jsonFile, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	err = json.Unmarshal(jsonFile, &tasks)

	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	return tasks
}
