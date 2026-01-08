package tasks_test

import (
	"taskmanager/tasks"
	"testing"
)

//Test the length method of IdSet by adding several IDs to the ta
func TestIdSetAdd(t *testing.T) {
	idSet := tasks.NewIdSet()
	idSet.Add(1)
	idSet.Add(2)
	if idSet.Length() != 2 {
		t.Errorf("Expected length to be 2, got %d", idSet.Length())
	}
}

//Test the Remove method of IdSet to ensure the size is accurate after several adds and deletes.
func TestIdSetRemove(t *testing.T) {
	idSet := tasks.NewIdSet()
	idSet.Add(1)
	idSet.Add(2)
	idSet.Add(3)
	idSet.Remove(1)
	idSet.Remove(3)
	if idSet.Length() != 1 {
		t.Errorf("Expected length to be 1 after removals, got %d", idSet.Length())
	}
}

//Test the Contains method of IdSet to ensure it correctly identifies existing and non-existing IDs.
func TestIdSetContains(t *testing.T) {
	idSet := tasks.NewIdSet()
	idSet.Add(1)
	idSet.Add(2)
	if !idSet.Contains(1) {
		t.Errorf("Expected ID 1 to be contained in the set")
	}
	if idSet.Contains(3) {
		t.Errorf("Expected ID 3 to not be contained in the set")
	}
}