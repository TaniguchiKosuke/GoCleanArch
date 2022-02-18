package hanlder

import (
	"GoCleanArch/domain/model"
	"GoCleanArch/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoUsecase usecase.TodoUsecase
}

func NewTodoHandler(todoUsecase *usecase.TodoUsecase) *TodoHandler {
	todoHandler := TodoHandler{todoUsecase: todoUsecase}
	return &todoHandler
}

func (th *TodoHandler) ViewTodo(c *gin.Context) {
	models, err := th.todoUsecase.ViewTodo()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, models)
		return
	}

	c.JSON(http.StatusOK, models)
}

func (th *TodoHandler) SearchTodo(c *gin.Context) {
	text := c.Query("text")
	models, err := th.todoUsecase.SearchTodo(text)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, models)
		return
	}

	c.JSON(http.StatusOK, models)
}

func (th *TodoHandler) AddTodo(c *gin.Context) {
	var todo *model.Todo
	if err := c.Bind(&todo); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, todo)
		return
	}

	if err := th.todoUsecase.AddTodo(&todo); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Failed to add new Todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Successed in posting new Todo"})
}

func (th *TodoHandler) EditTodo(c *gin.Context) {
	var todo *model.Todo
	if err := c.Bind(&todo); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, todo)
		return
	}

	if err := th.todoUsecase.EditTodo(&todo); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, todo)
		return
	}

	c.JSON(http.StatusOK, todo)
}