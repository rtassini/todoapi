package domain

import (
	"errors"
	"fmt"
	"strings"
)

var ErrTodoInvalidInput = errors.New("todo invalid input")

type TodoStatus string

type Todo struct {
	ID        string
	Title     string
	Completed bool
}

func NewTodo(title string, completed bool) (Todo, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return Todo{}, fmt.Errorf("%w: title", ErrTodoInvalidInput)
	}

	return Todo{
		Title:     title,
		Completed: completed,
	}, nil
}

func (t Todo) Update(title string, completed bool) (Todo, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return Todo{}, fmt.Errorf("%w: title", ErrTodoInvalidInput)
	}

	t.Title = title
	t.Completed = completed
	return t, nil
}

func (t Todo) MarkAsCompleted() Todo {
	t.Completed = true
	return t
}
