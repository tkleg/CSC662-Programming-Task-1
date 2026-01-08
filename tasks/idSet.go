package tasks

/*
IdSet

Meant to mimic a set data structure for easy use elsewhere.
I chose this implementation as it was very simple to write and use.

The code for this was sourced originally from https://www.bytesizego.com/blog/set-in-golang, then modified.

Fields:
- ids: A map with integer keys representing task IDs and empty struct values (since we only care about the presence of the key, not any associated value).

Methods:
- NewIdSet(): A constructor function that initializes and returns a new IdSet with an empty map.
- Add(id int): Adds the specified ID to the set by inserting it into the map.
- Remove(id int): Removes the specified ID from the set by deleting it from the map.
- Contains(id int) bool: Checks if the specified ID exists in the set by looking it up in the map, returning true if found and false otherwise.
*/

type IdSet struct {
	ids map[int]struct{}
}

func NewIdSet() IdSet {
	return IdSet{
		ids: make(map[int]struct{}),
	}
}

func (s *IdSet) Add(id int) {
	s.ids[id] = struct{}{}
}

func (s *IdSet) Remove(id int) {
	delete(s.ids, id)
}

func (s IdSet) Contains(id int) bool {
	_, exists := s.ids[id]
	return exists
}

func (s IdSet) Length() int {
	return len(s.ids)
}