package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/rillmind/taskManagerCli/enum/status"
	"github.com/rillmind/taskManagerCli/src/task"
)

func main() {
	verbose := flag.Bool("v", false, "Activates verbose mode, displaying more operation details.")

	flag.Parse()

	args := flag.Args()

	comando := args[0]

	switch comando {
	case "add":
		if *verbose {
			fmt.Println("Adding task!")
		}

		newTask := task.CreateTask(args[1])
		fmt.Printf("Task add sucessfully (ID: %v)\n", newTask.ID)

	case "list":
		if *verbose {
			fmt.Println("Listando tarefas!")
		}

		if len(args) == 1 {
			task.ListTasks()
		} else {
			switch args[1] {
			case "done":
				task.ListTasksByStatus(status.DONE)

			case "todo":
				task.ListTasksByStatus(status.TODO)

			case "in-progress":
				task.ListTasksByStatus(status.IN_PROGRESS)
			}
		}

	case "update":
		if *verbose {
			fmt.Println("Updating task!")
		}

		id, _ := strconv.Atoi(args[1])
		task.UpdateTask(id, args[2])

	case "delete":
		if *verbose {
			fmt.Println("Deleting task!")
		}

		id, _ := strconv.Atoi(args[1])
		task.DeleteTask(id)

	case "mark-in-progress":
		if *verbose {
			fmt.Println("Marking task as in progress!")
		}

		id, _ := strconv.Atoi(args[1])
		task.Mark(id, status.IN_PROGRESS)

	case "mark-done":
		if *verbose {
			fmt.Println("Marking task as done!")
		}

		id, _ := strconv.Atoi(args[1])
		task.Mark(id, status.DONE)
	}
}
