package tasks

import "fmt"

/*
Task

Holds information about a single task, including its unique ID, description, and completion status.

Fields:
- Id: An integer representing the unique identifier for the task.
- Description: A string containing the details of the task.

Methods:
- Print(): Displays the task's ID, description, and status in a formatted manner.
- MarkDone(): Sets the status field to Done (Status enum value), marking the task as completed.
*/

type Task struct {
	Id          int
	Description string
	Status      Status
}

// Displays the task's ID, description, and status in a formatted manner.
func (t Task) Print() {
	fmt.Printf("%-9d | %-50s | %-10s\n", t.Id, t.Description, t.Status.String())
}

// Sets the status field to Done (Status enum value), marking the task as completed.
func (t *Task) MarkDone() {
	t.Status = Done
}