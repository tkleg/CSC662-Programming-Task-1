package main

import "fmt"

/*
Task

Holds information about a single task, including its unique ID, description, and completion status.

Fields:
- Id: An integer representing the unique identifier for the task.
- Description: A string containing the details of the task.
- Done: A boolean indicating whether the task is Pending (false) or Done (true).

Methods:
- print(): Displays the task's ID, description, and status in a formatted manner.
- markDone(): Sets the Done field to true, marking the task as completed.
*/
type Task struct {
	Id int
	Description string
	Done bool
}

// Displays the task's ID, description, and status in a formatted manner.
func (t Task) print() {
	status := "Pending"
	if t.Done {
		status = "Done"
	}
	fmt.Printf("%-9d | %-50s | %-10s\n", t.Id, t.Description, status)
}

// Sets the Done field to true, marking the task as completed.
func (t *Task) markDone() {
	t.Done = true
}