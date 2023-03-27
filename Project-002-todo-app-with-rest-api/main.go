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

	"github.com/gin-gonic/gin"
)

func getTodoAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Here goes the list of all todos")
}

func getTodo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Here goes the list of single todo"+string(c.Param("id")))
}

func main() {

	r := gin.New()
	r.GET("/todo", getTodoAll)
	r.GET("/todo/:id", getTodo)
	r.Run(":8086")

}
