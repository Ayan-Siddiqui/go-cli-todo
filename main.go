package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	cf := NewCommandFlags()
	storage.Load(&todos)

	fmt.Println("Welcome to the Todo CLI")
	fmt.Println("Type 'help' for commands. Type 'exit' to exit")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())

		if input == "exit" || input == "quit" {
			fmt.Println("Goodbye")
			break
		}

		if input == "help" {
			fmt.Println("🛠 Available Commands:")
			fmt.Println("  add <task>           Add a new todo with the given task")
			fmt.Println("  list                 List all todos with their status")
			fmt.Println("  delete <index>       Delete the todo at the given index (0-based)")
			fmt.Println("  toggle <index>       Toggle the completion status of a todo")
			fmt.Println("  edit <index> <task>  Edit the title of a todo at the given index")
			fmt.Println("  help                 Show this help message")
			fmt.Println("  exit / quit          Exit the application")
			continue
		}

		cf.Parse(input)
		cf.Execute(&todos)
		storage.Save(&todos)
	}

}
