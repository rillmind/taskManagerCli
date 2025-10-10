package main

import (
	"flag"
	"fmt"

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
			fmt.Print("Adicionando tarefa!\n")
		}
		task.CreateTask(args[1])
	}
}
