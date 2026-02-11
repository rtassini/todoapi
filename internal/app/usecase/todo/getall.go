package todo

import (
	"todoapi/internal/app/usecase"
	"todoapi/internal/domain"
	"todoapi/internal/infra/repository"
)

type (
	GetallOutput struct {
		ID        string
		Title     string
		Completed bool
	}
	Getall interface {
		Handle() ([]GetallOutput, error)
	}
	getall struct {
		repository repository.TodoRepository
	}
)

func NewGetall(repository repository.TodoRepository) *getall {
	return &getall{
		repository: repository,
	}
}

func (uc *getall) Handle() ([]GetallOutput, error) {
	todoList, err := uc.repository.GetAll()
	if err != nil {
		return []GetallOutput{}, usecase.NewError("fail to get all todos in the repository", err,
			usecase.ErrorTypeInternalError)
	}
	return uc.domainListToOutputList(todoList), nil
}

func (uc *getall) domainListToOutputList(todoList []domain.Todo) []GetallOutput {
	var outputList []GetallOutput
	var outputObj GetallOutput

	for _, todo := range todoList {
		outputObj = GetallOutput{
			ID:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
		}
		outputList = append(outputList, outputObj)
	}
	return outputList
}
