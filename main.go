package main

import (
	"bufio" // Get user input as a whole line
	"fmt"
	"os" // Access to stdIn
	"strings" // Use for trimming whitespace from user input
	"strconv" // Convert string input to int for taskIDs
)

func main(){
	fmt.Println("Welcome to the Task Manager!")

	// Initialize Data
	tasklist := Tasklist{
		Tasks: []Task{},
		HighestID: 0,
		IDs: newIdSet(),
	} // Empty Takslist to begin with, will be populated by user input
	var command string // Should be add, done, list, delete, help, or exit
	var taskDescription string // User-entered description for a new task
	var taskID int // Used to search for a task by ID
	
	// Create a scanner to read user input
	// fmt.Scanln not used to handle entire lines since user input may contain whitespace on the ends
	// This choice was made to be more flexible with user input
	scanner := bufio.NewReader(os.Stdin)

	for {
		
		// Get nexst command from user input and trim whitespace for switch statement
		fmt.Print("Enter command (add, done, list, delete, help, exit): ")
		command, _ = scanner.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
			case "add":
				fmt.Print("Enter task description: ")
				taskDescription = nextLineTrimmed(scanner)
				tasklist.addTask(taskDescription)
			case "done":
				fmt.Print("Enter task ID to mark as done: ")
				taskID = nextIntTrimmed(scanner)
				tasklist.markTaskDone(taskID)
			case "list":
				tasklist.print()
			case "help":
				fmt.Println("Available commands: add, done, list, delete, help, exit")
			case "delete":
				fmt.Print("Enter task ID to delete: ")
				taskID = nextIntTrimmed(scanner)
				tasklist.deleteTask(taskID)
			case "exit":
				fmt.Println("Exiting Task Manager. Goodbye!")
				return
			default:
				fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}// End of switch statement used to determine which command the user entered
	}// End of main loop asking for input and processing the command
}

// Retrieves the next line of user input and trims whitespace from the ends before returning it as a string.
// bufio.Reader is used to read the entire line of input, ensuring whitespace does not affect the string.
func nextLineTrimmed(scanner *bufio.Reader) string {
	line, _ := scanner.ReadString('\n')
	return strings.TrimSpace(line)
}

// Retrieves the next line of user input, trims whitespace from the ends, and converts it to an integer before returning it.
// bufio.Reader is used to read the entire line of input, ensuring whitespace does not affect the integer.
func nextIntTrimmed(scanner *bufio.Reader) int {
	line, _ := scanner.ReadString('\n')
	line = strings.TrimSpace(line)
	num, _ := strconv.Atoi(line)
	return num
}