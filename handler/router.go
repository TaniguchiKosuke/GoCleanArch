package handler

import (
	"github.com/gin-gonic/gin"
)

func InitRouting(r *gin.Engine, todoHandler TodoHandler) {
    r.GET("/", todoHandler.ViewTodo)
    r.GET("/search", todoHandler.SearchTodo)
    r.POST("/todoCreate", todoHandler.AddTodo)
    r.POST("/todoEdit", todoHandler.EditTodo)
}