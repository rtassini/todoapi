package main

import (
	"log"
	"todoapi/internal/app/usecase/todo"
	"todoapi/internal/infra/controller"
	"todoapi/internal/infra/db"
	"todoapi/internal/infra/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Print("Starting the server...")

	server := gin.Default()

	dbConnection, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer dbConnection.Close()

	TodoCreateRepository := repository.NewTodoRepository(dbConnection)
	TodoCreateUseCase := todo.NewCreate(TodoCreateRepository)
	TodoGetallUseCase := todo.NewGetall(TodoCreateRepository)
	TodoUpdateUseCase := todo.NewUpdate(TodoCreateRepository)
	TodoDeleteUseCase := todo.NewDelete(TodoCreateRepository)

	TodoController := controller.NewTodoController(TodoCreateUseCase, TodoGetallUseCase, TodoUpdateUseCase, TodoDeleteUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/todos", TodoController.Getall)
	server.POST("/todos", TodoController.Create)
	server.PUT("/todos/:id", TodoController.Update)
	server.DELETE("/todos/:id", TodoController.Delete)

	server.Run(":8080")
}
