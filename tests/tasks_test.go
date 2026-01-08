package tasks_test

import (
	"taskmanager/tasks"
	"testing"
)

// TestTaskMarkDone verifies that the MarkDone method correctly updates a task's status to Done.
func TestTaskMarkDone(t *testing.T) {
	task := tasks.Task{
		Id: 1,
		Description: "Test task",
		Status: tasks.Pending,
	}
	task.MarkDone()
	if task.Status != tasks.Done {
		t.Errorf("Expected task status to be Done, got %s", task.Status.String())
	}
}

func TestTaskListAdd(t *testing.T) {
	tasklist := tasks.Tasklist{
		Tasks: []tasks.Task{},
		HighestID: 0,
		IDs: tasks.NewIdSet(),
	}
	tasklist.AddTask("Test task 1")
	tasklist.AddTask("Test task 2")
	if len(tasklist.Tasks) != 2 {
		t.Errorf("Expected 2 tasks in the list, got %d", len(tasklist.Tasks))
	}
	if tasklist.Tasks[0].Description != "Test task 1" || tasklist.Tasks[1].Description != "Test task 2" {
		t.Errorf("Task descriptions do not match expected values")
	}
}

func TestTaskListMarkDone(t *testing.T) {
	tasklist := tasks.Tasklist{
		Tasks: []tasks.Task{},
		HighestID: 0,
		IDs: tasks.NewIdSet(),
	}
	tasklist.AddTask("Test task 1")
	tasklist.AddTask("Test task 2")
	tasklist.MarkTaskDone(1)
	if tasklist.Tasks[0].Status != tasks.Done {
		t.Errorf("Expected task with ID 1 to be marked as Done, got %s", tasklist.Tasks[0].Status.String())
	}
	if tasklist.Tasks[1].Status != tasks.Pending {
		t.Errorf("Expected task with ID 2 to still be Pending, got %s", tasklist.Tasks[1].Status.String())
	}
}

func TestTaskListDelete(t *testing.T) {
	tasklist := tasks.Tasklist{
		Tasks: []tasks.Task{},
		HighestID: 0,
		IDs: tasks.NewIdSet(),
	}
	tasklist.AddTask("Test task 1")
	tasklist.AddTask("Test task 2")
	tasklist.DeleteTask(1)
	if len(tasklist.Tasks) != 1 {
		t.Errorf("Expected 1 task in the list after deletion, got %d", len(tasklist.Tasks))
	}
	if tasklist.Tasks[0].Id != 2 {
		t.Errorf("Expected remaining task to have ID 2, got %d", tasklist.Tasks[0].Id)
	}
}