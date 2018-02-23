package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Task struct {
	Id int
	// ...
}

// to call: GetById(1)
func GetById(id int) (result Task, err error) {
	// Todo
}

// to call: taskInstance.Update(...)
func (instance *Task) Update() (err error) {
	// Todo
}
