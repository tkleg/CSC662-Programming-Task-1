package main

import (
	"fmt"
	"strings"
	"slices"
)

/*
TaskList

Purpose: Manages a list of tasks, allowing for adding, marking as done, deleting, and printing tasks.

Fields:
- Tasks: A list of Task structs representing the current tasks in the list.
- HighestID: An integer tracking the highest assigned taskID to ensure auto-incrementing IDs for new tasks. Should start at 0 since the first task will be assigned taskID 1.
- IDs: An idSet struct to efficiently track existing task IDs for quick lookup when marking tasks as done or deleting them.

Methods:
- print(): Displays all tasks in a formatted table, showing their ID, description, and status (done or pending).
- addTask(description string): Creates a new task with a unique ID and the provided description, adds it to the task list, and updates the ID tracking.
- markTaskDone(id int): Marks the task with the specified ID as done if it exists, otherwise prints an error message.
- deleteTask(id int): Deletes the task with the specified ID if it exists, otherwise prints an error message.
*/

type Tasklist struct {
	Tasks []Task
	HighestID int
	IDs idSet
}

// Displays all tasks in a formatted table, showing all fields.
// If no tasks exist, a message is instead printed.
func (tl Tasklist) print() {
	if len(tl.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println()
	barrierLen, _ := fmt.Printf("%-9s | %-50s | %-10s\n", "ID", "Description", "Status")
	fmt.Println(strings.Repeat("-", barrierLen))
	for _, t := range tl.Tasks {
		t.print()
	}// Finds the tasks with the specified taskID
	fmt.Println()
}

// Increments the HighestID to ensure a unique ID for the new task.
// Adds a new task with the specified description to the Tasks list.
// Adds the new taskID to the IDs set for tracking.
func (tl *Tasklist) addTask(description string) {
	tl.HighestID++
	newTask := Task{
		Id: tl.HighestID,
		Description: description,
		Status: Pending,
	}
	tl.Tasks = append(tl.Tasks, newTask)
	tl.IDs.add(newTask.Id)
}

// Marks the task with the specified ID as done if it exists.
// If the task ID does not exist, an error message is printed.
func (tl *Tasklist) markTaskDone(id int) {
	if tl.IDs.contains(id) {	
		for i, t := range tl.Tasks {
			if t.Id == id {
				tl.Tasks[i].markDone()
				fmt.Printf("Task with ID %d marked as done.\n", id)
				return
			}
		}
	} else {
		fmt.Printf("Task with ID %d not found. Cannot mark done.\n", id)
	}
}

// Deletes the task with the specified ID if it exists.
// If the task ID does not exist, an error message is printed.
func (tl *Tasklist) deleteTask(id int) {
	if tl.IDs.contains(id) {
		for i, t := range tl.Tasks {
			if t.Id == id {
				tl.Tasks = slices.Delete(tl.Tasks, i, i+1)
				tl.IDs.remove(id)
				fmt.Printf("Task with ID %d deleted.\n", id)
				return
			}
		}
	} else {
		fmt.Printf("Task with ID %d not found. Cannot delete.\n", id)
	}
}