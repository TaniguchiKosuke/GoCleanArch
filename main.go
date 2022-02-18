package main

import (
	"GoCleanArch/handler"
	"GoCleanArch/injector"

	"github.com/gin-gonic/gin"
)

func main() {
	todoHandler := injector.InjectTodoHandler()
	r := gin.Default()
	handler.InitRouting(r, todoHandler)
	r.Run(":8000")
}