package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}
func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("not found ")
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
	}
	for i := range len(todos) {
		if todos[i].ID == id {
			context.IndentedJSON(http.StatusOK, gin.H{"Item deleted": todo})
			todos = append(todos[i+1:], todos[:i]...)
			return
		}
	}

}

func getTodo(context *gin.Context) {
	//context.IndentedJSON(http.StatusOK, todos)
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}
func togleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)

}

func postTodo(context *gin.Context) {
	var newTodo todo
	err := context.BindJSON(&newTodo)
	if err != nil {
		return
	}
	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}
func main() {
	router := gin.Default()
	router.GET("/todo", getTodos)
	router.GET("/todo/:id", getTodo)
	router.DELETE("/todo/:id", deleteTodo)
	router.PATCH("/todo/:id", togleTodoStatus)
	router.POST("/todo", postTodo)
	router.Run("localhost:8080")
}
