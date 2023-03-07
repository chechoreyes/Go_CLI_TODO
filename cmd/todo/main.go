package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/chechoreyes/cli_todoapp/cmd"
)

const (
	todoFile = ".todos.json"
)

func main() {

	// FLAGS in the terminal

	// Add
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	// Completed

	flag.Parse()

	todos := &todo.Todos{}

	// Shortcut of if, err is declared only in scope of if
	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.Add("Sample todo")
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(0)
		}
	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	default:
		fmt.Fprintf(os.Stdout, "invalid command")
	}

}
