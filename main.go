package main

import "fmt"

func main(){
	fmt.Println("Welcome to the Task Manager!")
	tasklist := Tasklist{
		Tasks: []Task{},
		HighestID: 0,
		CurrentIDs: make(map[int]struct{}),
	}

	tasklist.AddTask("Buy groceries")
	tasklist.AddTask("Finish homework")
	tasklist.AddTask("Call mom")
	PrintTasks(&tasklist)

	tasklist.MarkTaskDone(2)
	PrintTasks(&tasklist)
	
}