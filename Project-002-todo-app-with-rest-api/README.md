# TODO Application

* A simple todo application with the RESTful api using the gin.
* Goal is to have the following CRUD functionality along with rest api support

## Supported CRUD operations

- [x] GET todo
- [x] GET ALL todo

<br>

- [x] CREATE todo
- [x] CREATE multiple todo

<br>

- [x] UPDATE todo

<br>

- [x] DELETE todo
- [x] DELETE ALL todo




## Learnings

1. variables name inside struct has to start with caps else it wont be accessible outside. I was getting empty constantly in the `c.IndentedJSON(http.StatusOK, completeTodo)`
2. `ShouldBindJSON` should be sent with the address of the var ie `&varName`, else it was always failed to create/update in the `ShouldBindJson`.


## References

* https://youtu.be/d_L64KT3SFM
* https://go.dev/doc/tutorial/web-service-gin
* https://www.java67.com/2022/12/projects-you-can-build-to-learn-golang.html Point 5
* https://youtu.be/OoNeWiJ1Ebk?t=266
* https://stackoverflow.com/a/59509628/6222977
