package model

import (
	"gorm.io/gorm"
)

type Todo struct {
	*gorm.Model
    ID        int    `json:"id"`
    Task      string `json:"task"`
    LimitDate string `json:"limit_date"`
    Status    bool   `json:"status"`
}