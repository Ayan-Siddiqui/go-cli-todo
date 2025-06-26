package main

import (
	"fmt"
	"strconv"
	"strings"
)

type CommandFlags struct {
	Title   string
	Index   int
	Command string
}

func NewCommandFlags() *CommandFlags {
	return &CommandFlags{}
}

func (cf *CommandFlags) Parse(input string) {
	args := strings.Fields(input)
	if len(args) == 0 {
		return
	}

	cf.Command = args[0]

	switch cf.Command {
	case "add":
		if len(args) > 1 {
			cf.Title = strings.Join(args[1:], " ")
		}
		fmt.Printf("Todo added: %+v\n", cf.Title)

	case "edit":
		if len(args) > 2 {
			cf.Index = parseIndex(args[1])
			cf.Title = strings.Join(args[2:], " ")
		}
		fmt.Printf("Todo with id %+v edited\n", cf.Index)

	case "delete":
		if len(args) > 1 {
			cf.Index = parseIndex(args[1])
		}
		fmt.Printf("Todo with id %+v deleted\n", cf.Index)

	case "toggle":
		if len(args) > 1 {
			cf.Index = parseIndex(args[1])
		}
		fmt.Printf("Completed status changed\n")

	default:
		fmt.Printf("%+v is not a valid command\n", cf.Command)
	}
}

func (cf *CommandFlags) Execute(todos *Todos) {
	switch cf.Command {

	case "list":
		todos.print()

	case "add":
		todos.add(cf.Title)

	case "edit":
		todos.edit(cf.Index, cf.Title)

	case "delete":
		todos.delete(cf.Index)

	case "toggle":
		todos.toggle(cf.Index)
	}
}

func parseIndex(s string) int {
	index, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return index
}
