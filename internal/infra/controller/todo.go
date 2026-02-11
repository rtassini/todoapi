package controller

import (
	"todoapi/internal/app/usecase/todo"

	"github.com/gin-gonic/gin"
)

type (
	todoController struct {
		create todo.Create
		update todo.Update
		getall todo.Getall
		delete todo.Delete
	}
)

func NewTodoController(create todo.Create, getall todo.Getall, update todo.Update, delete todo.Delete) *todoController {
	return &todoController{
		create: create,
		getall: getall,
		update: update,
		delete: delete,
	}
}

func (h *todoController) Create(ctx *gin.Context) {
	var createInput todo.CreateInput
	err := ctx.BindJSON(&createInput)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	createOutput, err := h.create.Handle(createInput)

	ctx.JSON(200, gin.H{"response": createOutput})
}

func (h *todoController) Getall(ctx *gin.Context) {

	getallOutputList, err := h.getall.Handle()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"response": getallOutputList})
}

func (h *todoController) Update(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	var updateInput todo.UpdateInput
	err := ctx.BindJSON(&updateInput)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updateOutput, err := h.update.Handle(updateInput, id)

	ctx.JSON(200, gin.H{"response": updateOutput})
}

func (h *todoController) Delete(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	deleteOutput, err := h.delete.Handle(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"response": deleteOutput})
}
