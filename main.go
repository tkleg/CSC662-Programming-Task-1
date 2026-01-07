package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main(){
	fmt.Println("Welcome to the Task Manager!")
	tasklist := Tasklist{
		Tasks: []Task{},
		HighestID: 0,
		CurrentIDs: make(map[int]struct{}),
	}
	scanner := bufio.NewReader(os.Stdin)
	var command string
	var taskDescription string
	var taskID int
	var taskIDStr string
	var err error
	_ = err
	for {
		fmt.Print("Enter command (add, done, list, delete, help, exit): ")
		command, err = scanner.ReadString('\n')
		command = command[:len(command)-1]
		command = strings.TrimSpace(command)
		switch command {
			case "add":
				fmt.Print("Enter task description: ")
				taskDescription, err = scanner.ReadString('\n')
				taskDescription = strings.TrimSpace(taskDescription)
				tasklist.AddTask(taskDescription)
				fmt.Println("Task added successfully!")
			case "done":
				fmt.Print("Enter task ID to mark as done: ")
				taskIDStr, err = scanner.ReadString('\n')
				taskIDStr = strings.TrimSpace(taskIDStr)
				taskID, err = strconv.Atoi(taskIDStr)
				tasklist.MarkTaskDone(taskID)
			case "list":
				tasklist.print()
			case "help":
				fmt.Println("Available commands: add, done, list, delete, help, exit")
			case "delete":
				fmt.Print("Enter task ID to delete: ")
				taskIDStr, err = scanner.ReadString('\n')
				taskIDStr = strings.TrimSpace(taskIDStr)
				taskID, err = strconv.Atoi(taskIDStr)
				tasklist.DeleteTask(taskID)
			case "exit":
				fmt.Println("Exiting Task Manager. Goodbye!")
				return
			default:
				fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}