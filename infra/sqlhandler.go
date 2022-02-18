package infra

import (
	"GoCleanArch/domain/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqlHandler struct {
	Conn *gorm.DB
}

func NewSqlHandler() *SqlHandler {
	conn, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil
	}
	conn.AutoMigrate(&model.Todo{})

	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}