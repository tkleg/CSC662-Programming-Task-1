package main

import "fmt"

type Task struct {
	Id int
	Description string
	Done bool
}

type Tasklist struct {
	Tasks []Task
	HighestID int
	CurrentIDs map[int]struct{}
}

func printTask(t Task) {
	status := "Pending"
	if t.Done {
		status = "Done"
	}
	fmt.Printf("ID: %d | Description: %s | Status: %s\n", t.Id, t.Description, status)
}

func PrintTasks(tl *Tasklist) {	if len(tl.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, t := range tl.Tasks {
		printTask(t)
	}
	//Print currentTasksMap
	fmt.Println("Current Task IDs:")
	for id := range tl.CurrentIDs {
		fmt.Printf("ID: %d\n", id)
	}
	fmt.Println("Highest ID:", tl.HighestID)
}

func (tl *Tasklist) AddTask(description string) {
	tl.HighestID++
	newTask := Task{
		Id: tl.HighestID,
		Description: description,
		Done: false,
	}
	tl.Tasks = append(tl.Tasks, newTask)
	tl.CurrentIDs[newTask.Id] = struct{}{}
}

func (tl *Tasklist) MarkTaskDone(id int) {
	if _, exists := tl.CurrentIDs[id]; !exists {
		fmt.Printf("Task with ID %d not found. Cannot mark done.\n", id)
		return
	}
	for i, t := range tl.Tasks {
		if t.Id == id {
			tl.Tasks[i].Done = true
			return
		}
	}
}

func (tl *Tasklist) DeleteTask(id int) {
	if _, exists := tl.CurrentIDs[id]; !exists {
		fmt.Printf("Task with ID %d not found. Cannot delete.\n", id)
		return
	}
	for i, t := range tl.Tasks {
		if t.Id == id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			delete(tl.CurrentIDs, id)
			return
		}
	}
}