package handler

import (
	"go-todo-app/model"
	"go-todo-app/view/components"
	"go-todo-app/view/index"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	Todos []model.Todo
}

// SHOW INDEX ON INITIAL PAGE LOAD
func (h *TodoHandler) ShowIndex(c echo.Context) error {
	return render(c, index.Index(h.Todos))
}

// UPDATE TODO LIST WHEN TODO IS ADDED
func (h *TodoHandler) AddTodo(c echo.Context) error {
	text := c.FormValue("text")
	id := 1
	if len(h.Todos) > 0 {
		id = h.Todos[len(h.Todos)-1].ID + 1
	}
	todo := model.Todo{ID: id, Text: text, Complete: false}
	h.Todos = append(h.Todos, todo)
	return render(c, components.TodoList(h.Todos))
}

// UPDATE TODO LIST WHEN TODO IS REMOVED
func (h *TodoHandler) RemoveTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return render(c, components.TodoList(h.Todos))
	}
	for index, val := range h.Todos {
		if val.ID == id {
			h.Todos = append(h.Todos[:index], h.Todos[index+1:]...)
			break
		}
	}
	return render(c, components.TodoList(h.Todos))
}
