package main

import (
	"log"
	"os"

	"go-todo-app/handler"
	"go-todo-app/model"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	// SERVE THE CSS, INCLUDING THE TAILWIND CSS
	app.Static("/css", "css")

	// DATA TO START WITH FOR THE PRIMITIVE EXAMPLE
	todos := []model.Todo{
		{ID: 1, Text: "Do Dishes", Complete: false},
		{ID: 2, Text: "Clean Up Room", Complete: false},
	}
	handler := handler.TodoHandler{Todos: todos}

	// HANDLE INITIAL PAGE LOAD
	app.GET("/", handler.ShowIndex)

	// HANDLE ADD/REMOVE TODO REQUESTS
	app.POST("/add-todo", handler.AddTodo)
	app.POST("/remove-todo", handler.RemoveTodo)

	// CHECK FOR .ENV FILE
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// START SERVER
	app.Start(":" + os.Getenv("PORT"))
}
