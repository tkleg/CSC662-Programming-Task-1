# Module 1 Programming Task

## References
 - The set implementation was found [HERE](https://www.bytesizego.com/blog/set-in-golang).
 - A post by Miguel Pignatelli on May 7, 2012 found [HERE](https://groups.google.com/g/golang-nuts/c/OBsCksYHPG4) showed how to fix unused variable errors with `_ = varName`. This was only used during development and is not used in the final submitted version.
 - A method to handle enums with string representation included was found [HERE](https://gobyexample.com/enums).

## Overview
- This project serves to allow for simple management of some tasks entered by the user on the command line.
- The program provides the following options to the user
  - Adding tasks to the list
  - Removing tasks from the list
  - Marking tasks in the list as `Done`
  - Printing all tasks currently in the list in a table-formatted manner
  - Printing a help menu to show the user how to use the program
- This project utilizes the custom classes `IdSet`, `Task`, and `Tasklist` in addition to an enum named `Status`.
- Technically, they are structs, however the way the functions are called makes them operate more like instances of a class.
  - `Task`: A simple struct which holds an auto-generated ID, an user entered description, and a status that is either `Done` or `Pending`
    - This operates more closely to a struct since it only has a function for printing and one for changing the status from `Pending` to `Done`
  - `TaskList`: By far the most complex of the three classes. It utilizes a list of tasks, an integer tracking the highest taskID handed out, and a set of all taskIDs currently in the list.
    - the `Add` function of this class is what is used to simply add a task to the list.
  - `IdSet`: This is simply meant to serve as a way to implement a set. It utitilizes a map and related functions to mimic the adding of elements, the removal of elements, the search for a particular element, and the return of the size of the set. As shown above, this idea and code originated elsewhere. Credit to the author as I simply made small changes to fit my case.
  - `Status`: Simply enum that allows two options and to return a string representation of each enum value.

## Testing
- Unit tests are written for the three "classes" used in the program, with each function being tested by its testing function
- The tests are in a seperate package and folder for the purpose of keeping the directories smaller and easier to manage.
- To run the tests simply run `go test ./tests -v -count=1`
  - `-v` is optional, it lets **RUN** and **PASS** messages for each test to be printed along with the individual test runtimes.
  - `-count=1` is optional, but it ensures that cached tests are ignored.

## Build and Run
- To make an executable, run `go build main.go`.
  - After that, simply run `./main` to run the executable.
- To run without making an executable, run `go run main.go`