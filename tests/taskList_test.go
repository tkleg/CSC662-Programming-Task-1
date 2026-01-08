package tasks_test

import (
	"taskmanager/tasks"
	"testing"
)

// Adds 3 tests for the addTask method of Tasklist.
// Verifies the number of tasks in the list, the descriptions of the tasks, and the length of the taskIDs set after adding tasks.
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
	if tasklist.IDs.Length() != 2 {
		t.Errorf("Expected 2 IDs in the set, got %d", tasklist.IDs.Length())
	}
}

// Tests for the markTaskDone method of Tasklist.
// Verifies that the correct task is marked as done and that other tasks remain pending.
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

// Tests for the deleteTask method of Tasklist.
// Verifies the correct task is deleted, the correct one remains, and that the IDs set is updated accordingly.
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
	if tasklist.IDs.Length() != 1 || !tasklist.IDs.Contains(2) || tasklist.IDs.Contains(1) {
		t.Errorf("Expected only the ID 2 to be in the set after deletion, got length %d, contains 2: %t, contains 1: %t", tasklist.IDs.Length(), tasklist.IDs.Contains(2), tasklist.IDs.Contains(1))
	}
}