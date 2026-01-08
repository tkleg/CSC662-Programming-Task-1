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