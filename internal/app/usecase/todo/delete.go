package todo

import (
	"fmt"
	"todoapi/internal/app/usecase"
	"todoapi/internal/infra/repository"
)

type (
	DeleteOutput struct {
		Message string
	}
	Delete interface {
		Handle(string) (DeleteOutput, error)
	}
	delete struct {
		repository repository.TodoRepository
	}
)

func NewDelete(repository repository.TodoRepository) *delete {
	return &delete{
		repository: repository,
	}
}

func (uc *delete) Handle(id string) (DeleteOutput, error) {
	rowsUpdated, err := uc.repository.DeleteByID(id)
	if err != nil {
		return DeleteOutput{}, usecase.NewError("fail to delete a todo in the repository", err,
			usecase.ErrorTypeInternalError)
	}
	return uc.domainTodoToDeleteOutput(rowsUpdated), nil
}

func (uc *delete) domainTodoToDeleteOutput(rowsUpdated int64) DeleteOutput {
	return DeleteOutput{
		Message: fmt.Sprint("Rows deleted: ", rowsUpdated, ""),
	}
}
