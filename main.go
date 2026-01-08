package main

import (
	"bufio" // Get user input as a whole line
	"fmt"
	"os" // Access to stdIn
	"strings" // Use for trimming whitespace from user input
	"strconv" // Convert string input to int for taskIDs
	"taskmanager/tasks"
)

const HELP_MESSAGE string = `
Available commands:
	- add: Add a new task with a description. You will be prompted to enter a description after entering this command.
	- done: Mark a task as done by providing its ID. You will be prompted to enter the task ID after entering this command.
	- list: List all tasks with their ID, description, and status.
	- delete: Delete a task by providing its ID. You will be prompted to enter the task ID after entering this command.
	- help: Display this help message.
	- exit: Exit the Task Manager.
`

func main(){
	fmt.Println("Welcome to the Task Manager!")

	// Initialize Data
	tasklist := tasks.Tasklist{
		Tasks: []tasks.Task{},
		HighestID: 0,
		IDs: tasks.NewIdSet(),
	} // Empty Tasklist to begin with, will be populated by user input
	var command string // Should be add, done, list, delete, help, or exit
	var taskDescription string // User-entered description for a new task
	var taskID int // Used to search for a task by ID
	var err error // Used to check for errors when user is prompted to enter a command

	// Create a scanner to read user input
	// fmt.Scanln not used to handle entire lines since user input may contain whitespace on the ends
	// This choice was made to be more flexible with user input
	scanner := bufio.NewReader(os.Stdin)

	for {
		
		// Get nexst command from user input and trim whitespace for switch statement
		command, err = getNextCommand(scanner)
		if err != nil {
			fmt.Println("Error reading command. Please try again.")
			continue
		}

		switch command {
			case "add":
				fmt.Print("Enter task description: ")
				taskDescription, err = nextLineTrimmed(scanner)
				if err == nil {
					tasklist.AddTask(taskDescription)
				}// No print statement for error case since nextLineTrimmed already handles that
			case "done":
				fmt.Print("Enter task ID to mark as done: ")
				taskID, err = nextIntTrimmed(scanner)
				if err == nil && taskID > 0 { // Valid integer and greater than 0
					tasklist.MarkTaskDone(taskID)
				} else if err == nil { // Valid integer but not greater than 0
					fmt.Println("Invalid task ID. Please enter an integer greater than 0. The entered value was:", taskID)
				}
			case "list":
				tasklist.Print()
			case "help":
				fmt.Println(HELP_MESSAGE)
			case "delete":
				fmt.Print("Enter task ID to delete: ")
				taskID, err = nextIntTrimmed(scanner)
				if err == nil && taskID > 0 { // Valid integer and greater than 0
					tasklist.DeleteTask(taskID)
				} else if err == nil { // Valid integer but not greater than 0
					fmt.Println("Invalid task ID. Please enter an integer greater than 0. The entered value was:", taskID)
				}
			case "exit":
				fmt.Println("Exiting Task Manager. Goodbye!")
				return
			default:
				fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}// End of switch statement used to determine which command the user entered
	}// End of main loop asking for input and processing the command
}

func getNextCommand(scanner *bufio.Reader) (string, error) {
	fmt.Print("Enter command (add, done, list, delete, help, exit): ")
	command, err := scanner.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(command), nil
}

// Retrieves the next line of user input and trims whitespace from the ends before returning it as a string.
// bufio.Reader is used to read the entire line of input, ensuring whitespace does not affect the string.
func nextLineTrimmed(scanner *bufio.Reader) (string, error) {
	line, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input. Please enter a valid string.")
		return "\x00", err
	}
	return strings.TrimSpace(line), nil
}

// Retrieves the next line of user input, trims whitespace from the ends, and converts it to an integer before returning it.
// bufio.Reader is used to read the entire line of input, ensuring whitespace does not affect the integer.
func nextIntTrimmed(scanner *bufio.Reader) (int, error) {

	// Read the next line of input and trim whitespace
	line, err := scanner.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input. Please enter a valid integer.")
		return -1, err
	}
	line = strings.TrimSpace(line)

	// Convert the trimmed string to an integer
	num, err := strconv.Atoi(line)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid integer.")
		return -1, err
	}

	return num, nil
}