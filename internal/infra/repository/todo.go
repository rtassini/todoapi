package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"todoapi/internal/domain"
)

type TodoRepository struct {
	connection *sql.DB
}

func NewTodoRepository(connection *sql.DB) TodoRepository {
	return TodoRepository{connection: connection}
}

func (c *TodoRepository) Create(title string, completed bool) (string, error) {

	var id []uint8
	query, err := c.connection.Prepare("INSERT INTO todo (title, completed) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	err = query.QueryRow(title, completed).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(id), nil

}

func (c *TodoRepository) DeleteByID(id string) (int64, error) {
	query := "DELETE FROM todo WHERE id = $1"

	result, err := c.connection.Exec(query, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return 0, errors.New("no record found with that ID")
	}

	return rowsAffected, nil
}

func (c *TodoRepository) GetAll() ([]domain.Todo, error) {

	query := "select id, title, completed FROM todo"
	rows, err := c.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []domain.Todo{}, err
	}
	var todoList []domain.Todo
	var todoObj domain.Todo
	for rows.Next() {
		err := rows.Scan(&todoObj.ID, &todoObj.Title, &todoObj.Completed)
		if err != nil {
			fmt.Println(err)
			return []domain.Todo{}, err
		}
		todoList = append(todoList, todoObj)
	}

	rows.Close()
	return todoList, nil
}

func (c *TodoRepository) UpdateByID(id, title string, completed bool) (int64, error) {

	query := "UPDATE todo SET title = $1, completed = $2 WHERE id = $3"

	result, err := c.connection.Exec(query, title, completed, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return 0, errors.New("no record found with that ID")
	}

	return rowsAffected, nil
}
