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
			fmt.Println("Adicionando tarefa!")
		}
		task.CreateTask(args[1])

	case "listAll":
		if *verbose {
			fmt.Println("Listando tarefas!")
		}
		task.ListTasks()

	case "delete":
		if *verbose {
			fmt.Println("Removendo tarefa")
		}
		id, _ := strconv.Atoi(args[1])
		task.DeleteTask(id)
	}
}
