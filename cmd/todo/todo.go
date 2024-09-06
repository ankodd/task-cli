package todo

import (
	"time"
)

type Todo struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ID          int
	State       State
	Description string
}

func NewTodo(id int, createdAt time.Time, description string) Todo {
	return Todo{
		CreatedAt:   createdAt,
		UpdatedAt:   time.Now(),
		ID:          id,
		State:       Added,
		Description: description,
	}
}

func (t *Todo) SetState(st State) {
	t.State = st
}
