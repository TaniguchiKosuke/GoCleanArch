package injector

import (
	"GoCleanArch/domain/repository"
	"GoCleanArch/infra"
	"GoCleanArch/usecase"
	"GoCleanArch/handler"
)

func InjectDB() infra.SqlHandler {
    sqlhandler := infra.NewSqlHandler()
    return *sqlhandler
}

func InjectTodoRepository() repository.TodoRepository {
    sqlHandler := InjectDB()
    return infra.NewTodoRepository(&sqlHandler)
}

func InjectTodoUsecase() usecase.TodoUsecase {
    TodoRepo := InjectTodoRepository()
    return usecase.NewTodoUsecase(TodoRepo)
}

func InjectTodoHandler() handler.TodoHandler {
    return handler.NewTodoHandler(InjectTodoUsecase())
}