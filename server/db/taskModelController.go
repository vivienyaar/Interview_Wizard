package db

import (
// "database/sql"
// "fmt"
// _ "github.com/lib/pq"
)

type Task struct {
	Id int
	// ...
}

// to call: GetFirstUnfinished()
func GetFirstUnfinished() (result Task, err error) {
	// Todo
	return Task{}, nil
}

// to call: taskInstance.Update(id, ...)
func (instance *Task) Update(id int) (err error) {
	// Todo
	return nil
}
