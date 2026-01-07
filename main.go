package main

import "fmt"

func main(){
	fmt.Println("Welcome to the Task Manager!")
	tasklist := Tasklist{
		Tasks: []Task{},
		HighestID: 0,
		CurrentIDs: make(map[int]struct{}),
	}
	_ = tasklist // Suppress unused variable error
	var command string
	for {
		fmt.Print("Enter command (add, done, list, help, exit): ")
		fmt.Scan(&command)
		switch command {
			case "add":
				fmt.Println("Adding")
			case "done":
				fmt.Println("Marking done")
			case "list":
				fmt.Println("Listing tasks")
			case "help":
				fmt.Println("Available commands: add, done, list, help, exit")
			case "exit":
				fmt.Println("Exiting Task Manager. Goodbye!")
				return
			default:
				fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}