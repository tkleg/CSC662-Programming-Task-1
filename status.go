package main

/*
Status

Defines possible statuses for a task and provides a mapping to their string representations for display purposes.

This method of enum handling was found on https://gobyexample.com/enums.

Values: Done and Pending

Fields:
- statusName: A map that maps each Status value to its corresponding string representation ("Done" or "Pending").

Methods:
- String() string: Returns the string representaton of the calling Status value

*/
type Status int
const (
	Done Status = iota
	Pending
)

var statusName = map[Status]string{
	Done: "Done",
	Pending: "Pending",
}

func (s Status) String() string {
	return statusName[s]
}