package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTodo(t *testing.T) {
	//exampleTitle := "title example"
	//exampleTodo := Todo{
	//	Title:     exampleTitle,
	//	Completed: false,
	//}
	testCases := []struct {
		name      string
		title     string
		completed bool
		result    Todo
		err       error
	}{
		{
			name:   "should fail when title is invalid",
			title:  "",
			result: Todo{},
			err:    fmt.Errorf("%w: title", ErrTodoInvalidInput),
		},
		//{
		//	name:        "should fail when title with space is invalid",
		//	title:       " ",
		//	description: exampleDescription,
		//	date:        exampleDate,
		//	dueDate:     nil,
		//	result:      domain.Todo{},
		//	err:         fmt.Errorf("%w: title", domain.ErrTodoInvalidInput),
		//},
		//{
		//	name:        "should fail when date is invalid",
		//	title:       exampleTitle,
		//	description: exampleDescription,
		//	date:        time.Time{},
		//	dueDate:     nil,
		//	result:      domain.Todo{},
		//	err:         fmt.Errorf("%w: date", domain.ErrTodoInvalidInput),
		//},
		//{
		//	name:        "should create todo",
		//	title:       exampleTitle,
		//	description: exampleDescription,
		//	date:        exampleDate,
		//	dueDate:     nil,
		//	result:      exampleTodo,
		//	err:         nil,
		//},
		//{
		//	name:        "should create todo with title and description with spaces",
		//	title:       " title example ",
		//	description: " description example ",
		//	date:        exampleDate,
		//	dueDate:     nil,
		//	result:      exampleTodo,
		//	err:         nil,
		//},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := NewTodo(tc.title, false)
			assert.Equal(t, tc.result, result)
			assert.Equal(t, tc.err, err)
		})
	}
}

//func TestTodo_Update(t *testing.T) {
//	exampleTitle := "title example"
//	exampleTitleUpdated := "title example updated"
//	exampleDescription := "description example"
//	exampleDescriptionUpdated := "description example updated"
//	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
//	exampleDateUpdated, _ := time.Parse(time.DateOnly, "2024-01-02")
//	exampleTodo := domain.Todo{
//		Title:       exampleTitle,
//		Description: exampleDescription,
//		Status:      domain.TodoStatusPending,
//		DueDate:     nil,
//		CreatedAt:   exampleDate,
//		UpdatedAt:   exampleDate,
//	}
//	exampleTodoUpdated := domain.Todo{
//		Title:       exampleTitleUpdated,
//		Description: exampleDescriptionUpdated,
//		Status:      domain.TodoStatusPending,
//		DueDate:     nil,
//		CreatedAt:   exampleDate,
//		UpdatedAt:   exampleDateUpdated,
//	}
//	testCases := []struct {
//		name        string
//		title       string
//		description string
//		date        time.Time
//		dueDate     *time.Time
//		todo        domain.Todo
//		result      domain.Todo
//		err         error
//	}{
//		{
//			name:        "should fail when title is invalid",
//			title:       "",
//			description: exampleDescriptionUpdated,
//			date:        exampleDateUpdated,
//			dueDate:     nil,
//			todo:        exampleTodo,
//			result:      domain.Todo{},
//			err:         fmt.Errorf("%w: title", domain.ErrTodoInvalidInput),
//		},
//		{
//			name:        "should fail when title with space is invalid",
//			title:       " ",
//			description: exampleDescriptionUpdated,
//			date:        exampleDateUpdated,
//			dueDate:     nil,
//			todo:        exampleTodo,
//			result:      domain.Todo{},
//			err:         fmt.Errorf("%w: title", domain.ErrTodoInvalidInput),
//		},
//		{
//			name:        "should fail when date is invalid",
//			title:       exampleTitleUpdated,
//			description: exampleDescriptionUpdated,
//			date:        time.Time{},
//			dueDate:     nil,
//			todo:        exampleTodo,
//			result:      domain.Todo{},
//			err:         fmt.Errorf("%w: date", domain.ErrTodoInvalidInput),
//		},
//		{
//			name:        "should update todo",
//			title:       exampleTitleUpdated,
//			description: exampleDescriptionUpdated,
//			date:        exampleDateUpdated,
//			dueDate:     nil,
//			todo:        exampleTodo,
//			result:      exampleTodoUpdated,
//			err:         nil,
//		},
//		{
//			name:        "should update todo with title and description with spaces",
//			title:       " title example updated ",
//			description: " description example updated ",
//			date:        exampleDateUpdated,
//			dueDate:     nil,
//			todo:        exampleTodo,
//			result:      exampleTodoUpdated,
//			err:         nil,
//		},
//	}
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			result, err := tc.todo.Update(tc.title, tc.description, tc.date, tc.dueDate)
//			assert.Equal(t, tc.result, result)
//			assert.Equal(t, tc.err, err)
//		})
//	}
//}
//
//func TestTodo_MarkAsCompleted(t *testing.T) {
//	exampleTitle := "title example"
//	exampleDescription := "description example"
//	exampleDate, _ := time.Parse(time.DateOnly, "2024-01-01")
//	exampleDateUpdated, _ := time.Parse(time.DateOnly, "2024-01-02")
//	exampleTodo := domain.Todo{
//		Title:       exampleTitle,
//		Description: exampleDescription,
//		Status:      domain.TodoStatusPending,
//		DueDate:     nil,
//		CreatedAt:   exampleDate,
//		UpdatedAt:   exampleDate,
//	}
//	exampleTodoCompleted := domain.Todo{
//		Title:       exampleTitle,
//		Description: exampleDescription,
//		Status:      domain.TodoStatusCompleted,
//		DueDate:     nil,
//		CreatedAt:   exampleDate,
//		UpdatedAt:   exampleDateUpdated,
//	}
//	testCases := []struct {
//		name   string
//		todo   domain.Todo
//		date   time.Time
//		result domain.Todo
//	}{
//		{
//			name:   "should mark todo as completed",
//			todo:   exampleTodo,
//			date:   exampleDateUpdated,
//			result: exampleTodoCompleted,
//		},
//	}
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			result := tc.todo.MarkAsCompleted(tc.date)
//			assert.Equal(t, tc.result, result)
//		})
//	}
//}
