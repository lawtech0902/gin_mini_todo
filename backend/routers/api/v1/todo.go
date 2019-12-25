package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lawtech0902/gin_todo_demo/backend/models"
	"github.com/lawtech0902/gin_todo_demo/backend/pkg/app"
	"github.com/lawtech0902/gin_todo_demo/backend/pkg/e"
	"github.com/unknwon/com"
	"net/http"
)

type TodoList struct {
	Total uint64             `json:"total"`
	List  []models.TodoModel `json:"list"`
}

func FetchAllTodo(c *gin.Context) {
	var todo models.TodoModel
	
	count, todoList, err := todo.GetAll()
	if err != nil {
		app.SendResponse(c, http.StatusInternalServerError, e.ErrDatabase, nil)
		return
	}
	
	app.SendResponse(c, http.StatusOK, nil, TodoList{
		Total: count,
		List:  todoList,
	})
}

func FetchOneTodo(c *gin.Context) {
	var (
		todo models.TodoModel
		err  error
	)
	
	id := com.StrTo(c.Param("id")).MustInt()
	todo.ID = uint(id)
	
	if todo, err = todo.Get(); err != nil {
		app.SendResponse(c, http.StatusInternalServerError, e.ErrDatabase, nil)
		return
	}
	
	app.SendResponse(c, http.StatusOK, nil, todo)
}

func AddTodo(c *gin.Context) {
	completed := com.StrTo(c.PostForm("completed")).MustInt()
	
	todo := models.TodoModel{
		Title:     c.PostForm("title"),
		Completed: completed,
	}
	
	if err := todo.Create(); err != nil {
		app.SendResponse(c, http.StatusInternalServerError, e.ErrDatabase, nil)
		return
	}
	
	app.SendResponse(c, http.StatusCreated, nil, "create successful!")
}

func UpdateTodo(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	completed := com.StrTo(c.PostForm("completed")).MustInt()
	todo := models.TodoModel{
		Title:     c.PostForm("title"),
		Completed: completed,
	}
	
	todo.ID = uint(id)
	if err := todo.Update(); err != nil {
		app.SendResponse(c, http.StatusInternalServerError, e.ErrDatabase, nil)
		return
	}
	
	app.SendResponse(c, http.StatusOK, nil, "update successful!")
}

func DeleteTodo(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	todo := models.TodoModel{}
	
	todo.ID = uint(id)
	
	if err := todo.Delete(); err != nil {
		app.SendResponse(c, http.StatusInternalServerError, e.ErrDatabase, nil)
		return
	}
	
	app.SendResponse(c, http.StatusOK, nil, "delete successful!")
}
