package todo

import (
	"fmt"
	"todoapi/internal/app/usecase"
	"todoapi/internal/infra/repository"
)

type (
	UpdateInput struct {
		Title     string
		Completed bool
	}
	UpdateOutput struct {
		Message string
	}
	Update interface {
		Handle(input UpdateInput, id string) (UpdateOutput, error)
	}
	update struct {
		repository repository.TodoRepository
	}
)

func NewUpdate(repository repository.TodoRepository) *update {
	return &update{
		repository: repository,
	}
}

func (uc *update) Handle(input UpdateInput, id string) (UpdateOutput, error) {
	rowsUpdated, err := uc.repository.UpdateByID(id, input.Title, input.Completed)
	if err != nil {
		return UpdateOutput{}, usecase.NewError("fail to update a todo in the repository", err,
			usecase.ErrorTypeInternalError)
	}
	return uc.domainTodoToUpdateOutput(rowsUpdated), nil

}

func (uc *update) domainTodoToUpdateOutput(rowsUpdated int64) UpdateOutput {
	return UpdateOutput{
		Message: fmt.Sprint("Rows updated: ", rowsUpdated, ""),
	}
}
