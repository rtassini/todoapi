package todo

import (
	"todoapi/internal/app/usecase"
	"todoapi/internal/domain"
	"todoapi/internal/infra/repository"
)

type (
	CreateInput struct {
		Title     string
		Completed bool
	}
	CreateOutput struct {
		ID        string
		Title     string
		Completed bool
	}
	Create interface {
		Handle(input CreateInput) (CreateOutput, error)
	}
	create struct {
		store repository.TodoRepository
	}
)

func NewCreate(store repository.TodoRepository) *create {
	return &create{
		store: store,
	}
}

func (uc *create) Handle(input CreateInput) (CreateOutput, error) {
	todo, err := domain.NewTodo(input.Title, input.Completed)
	if err != nil {
		return CreateOutput{}, usecase.NewError(err.Error(), err, usecase.ErrorTypeBadRequest)
	}
	id, err := uc.store.Create(input.Title, input.Completed)
	if err != nil {
		return CreateOutput{}, usecase.NewError("fail to create a todo in the repository", err,
			usecase.ErrorTypeInternalError)
	}
	return uc.domainTodoToOutput(todo, id), nil
}

func (uc *create) domainTodoToOutput(todo domain.Todo, id string) CreateOutput {
	return CreateOutput{
		ID:        id,
		Title:     todo.Title,
		Completed: todo.Completed,
	}
}
