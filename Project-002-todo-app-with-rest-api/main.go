package main

/*

This is a simple TODO application

Goal is to have the following functionality along with rest api support

- GET single todo
- GET all todo
- CREATE single todo
- CREATE multple todo
- Update single todo to done-un-done
- Update all todo to done/un-done
- DELETE single todo
- DELETE all todo
*/

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var completeTodo = []todo{
	{1, "a1", false},
	{2, "a2", false},
}

func getTodoAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, completeTodo)
}

func getTodo(c *gin.Context) {
	todoId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err == nil {
		for _, td := range completeTodo {
			if td.Id == todoId {
				c.IndentedJSON(http.StatusOK, td)
				return
			}
		}
	}
	c.IndentedJSON(http.StatusNotFound, "Id not Found 404")
}

func delteTodoAll(c *gin.Context) {
	completeTodo = []todo{}
	c.IndentedJSON(http.StatusOK, completeTodo)
}

func delteTodo(c *gin.Context) {
	todoId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err == nil {
		for i, td := range completeTodo {
			if td.Id == todoId {
				completeTodo[i] = completeTodo[0]
				completeTodo = completeTodo[1:]

				c.IndentedJSON(http.StatusOK, "Deleted todo "+td.Name)
				return
			}
		}
	}
	c.IndentedJSON(http.StatusNotFound, "Id not Found to delete 404")
}

func main() {

	r := gin.New()
	r.GET("/todo", getTodoAll)
	r.GET("/todo/:id", getTodo)

	r.DELETE("/todo", delteTodoAll)
	r.DELETE("/todo/:id", delteTodo)

	r.Run(":8086")

}
