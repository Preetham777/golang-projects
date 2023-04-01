package main

/*
   .-----------------------.
   |C>_ Author: Preetham G |
   |                       |
 __|_______________________|__
|  __________________________--|
`-/.::::::::::::::::::::::::.\-'a
 `----------------------------'

*/

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
+---------------------------------------------------------+
| todo is a struct type                                   |
| Which holds Id, name and done                           |
| Make sure the beginning letter to be Capitalized        |
+---------------------------------------------------------+
*/

type todo struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

// Default 2 todo are added

var completeTodo = []todo{
	{1, "a1", false},
	{2, "a2", false},
}

/*
*********************************************************************
*  ██████╗ ███████╗████████╗                                        *
* ██╔════╝ ██╔════╝╚══██╔══╝                                        *
* ██║  ███╗█████╗     ██║                                           *
* ██║   ██║██╔══╝     ██║                                           *
* ╚██████╔╝███████╗   ██║                                           *
*  ╚═════╝ ╚══════╝   ╚═╝                                           *
*********************************************************************
 */
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

/*
***********************************************************************
* ██████╗ ███████╗██╗     ███████╗████████╗███████╗                   *
* ██╔══██╗██╔════╝██║     ██╔════╝╚══██╔══╝██╔════╝                   *
* ██║  ██║█████╗  ██║     █████╗     ██║   █████╗                     *
* ██║  ██║██╔══╝  ██║     ██╔══╝     ██║   ██╔══╝                     *
* ██████╔╝███████╗███████╗███████╗   ██║   ███████╗                   *
* ╚═════╝ ╚══════╝╚══════╝╚══════╝   ╚═╝   ╚══════╝                   *
***********************************************************************
 */
func deleteTodoAll(c *gin.Context) {
	completeTodo = []todo{}
	c.IndentedJSON(http.StatusOK, completeTodo)
}

func deleteTodo(c *gin.Context) {
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

/*
********************************************************************
 */

/*
***********************************************************************
*  ██████╗██████╗ ███████╗ █████╗ ████████╗███████╗                   *
* ██╔════╝██╔══██╗██╔════╝██╔══██╗╚══██╔══╝██╔════╝                   *
* ██║     ██████╔╝█████╗  ███████║   ██║   █████╗                     *
* ██║     ██╔══██╗██╔══╝  ██╔══██║   ██║   ██╔══╝                     *
* ╚██████╗██║  ██║███████╗██║  ██║   ██║   ███████╗                   *
*  ╚═════╝╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚══════╝                   *
***********************************************************************
 */

func createTodo(c *gin.Context) {
	var newTodo []todo

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Unable to create new todo")
		return
	}

	var allTodoNamesCreate string

	for _, ntd := range newTodo {
		completeTodo = append(completeTodo, ntd)
		allTodoNamesCreate = allTodoNamesCreate + " " + ntd.Name
		fmt.Printf("New todo to create %v\n", ntd)
	}

	c.IndentedJSON(http.StatusNotFound, "Todo "+allTodoNamesCreate+" created ")

}

/*
***********************************************************************
* ██╗   ██╗██████╗ ██████╗  █████╗ ████████╗███████╗                  *
* ██║   ██║██╔══██╗██╔══██╗██╔══██╗╚══██╔══╝██╔════╝                  *
* ██║   ██║██████╔╝██║  ██║███████║   ██║   █████╗                    *
* ██║   ██║██╔═══╝ ██║  ██║██╔══██║   ██║   ██╔══╝                    *
* ╚██████╔╝██║     ██████╔╝██║  ██║   ██║   ███████╗                  *
*  ╚═════╝ ╚═╝     ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝                  *
***********************************************************************
 */

func updateTodo(c *gin.Context) {
	var updateTodo todo

	todoId, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "Unable to update")
		return
	}

	if err := c.ShouldBindJSON(&updateTodo); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	for i, td := range completeTodo {
		if td.Id == todoId {
			completeTodo[i] = updateTodo
			c.IndentedJSON(http.StatusOK, "Updated todo "+updateTodo.Name)
			return
		}
	}

	c.IndentedJSON(http.StatusBadRequest, "Unable to update todo "+updateTodo.Name)

}

func main() {

	r := gin.New()

	// Get todo(s)
	r.GET("/todo", getTodoAll)
	r.GET("/todo/:id", getTodo)

	// Delete todo(s)
	r.DELETE("/todo", deleteTodoAll)
	r.DELETE("/todo/:id", deleteTodo)

	// Create todo(s)
	r.POST("/todo", createTodo)

	// Update todo
	r.PUT("/todo/:id", updateTodo)

	r.Run(":8086")

}
