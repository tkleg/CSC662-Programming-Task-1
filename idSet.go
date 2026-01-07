package main

/*
idSet

Meant to mimic a set data structure for easy use elsewhere

The code for this was sourced originally from https://www.bytesizego.com/blog/set-in-golang, then modified.

Fields:
- ids: A map with integer keys representing task IDs and empty struct values (since we only care about the presence of the key, not any associated value).

Methods:
- newIdSet(): A constructor function that initializes and returns a new idSet with an empty map.
- add(id int): Adds the specified ID to the set by inserting it into the map.
- remove(id int): Removes the specified ID from the set by deleting it from the map.
- contains(id int) bool: Checks if the specified ID exists in the set by looking it up in the map, returning true if found and false otherwise.
*/

type idSet struct {
	ids map[int]struct{}
}

func newIdSet() idSet {
	return idSet{
		ids: make(map[int]struct{}),
	}
}

func (s *idSet) add(id int) {
	s.ids[id] = struct{}{}
}

func (s *idSet) remove(id int) {
	delete(s.ids, id)
}

func (s idSet) contains(id int) bool {
	_, exists := s.ids[id]
	return exists
}