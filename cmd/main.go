package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/rillmind/taskManagerCli/enum/status"
	"github.com/rillmind/taskManagerCli/src/task"
)

func main() {
	verbose := flag.Bool("v", false, "Activates verbose mode, displaying more operation details")

	flag.Usage = printUsage

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		printUsage()
		os.Exit(1)
	}

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
	default:
		fmt.Printf("Command not recognized: %v!", args[0])
	}
}

func printUsage() {
	fmt.Printf(`
Task Manager CLI - Manage your tasks from the command line

Usage:
	task-cli [options] <command> [arguments]

Options:
	-v    Activate verbose mode, displaying more operation details

Commands:
	add <title>             Create a new task with the given title
	list                    List all tasks
	list done               List only completed tasks
	list todo               List only pending tasks
	list in-progress        List tasks in progress
	update <id> <title>     Update the title of a task
	delete <id>             Delete a task
	mark-in-progress <id>   Mark a task as in progress
	mark-done <id>          Mark a task as completed

Examples:
	task-cli add "Complete project documentation"
	task-cli list todo
	task-cli -v update 1 "Updated task title"
	task-cli mark-done 2
`)
}
