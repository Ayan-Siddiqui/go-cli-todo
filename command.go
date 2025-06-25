package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCommandFlags() *CommandFlags {
	cf := CommandFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo title by specifying its id. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo by specifying its id")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Change the 'completed' status of a todo")
	flag.BoolVar(&cf.List, "list", false, "List all the todos")

	flag.Parse()

	return &cf
}

func (cf *CommandFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()

	case cf.Add != "":
		todos.add(cf.Add)

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid id for edit")
			os.Exit(1)
		}
		todos.edit(index, parts[1])

	case cf.Del != -1:
		todos.delete(cf.Del)

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	}
}
