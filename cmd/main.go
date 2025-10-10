package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/rillmind/taskManagerCli/src/task"
)

func main() {
	verbose := flag.Bool("v", false, "Ativa o modo verbose, exibindo mais detalhes da operação.")

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
		task.ListTasks()

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
	}
}
