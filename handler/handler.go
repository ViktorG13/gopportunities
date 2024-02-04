package handler

import (
	"github.com/ViktorG13/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializerHandle() {
	db = config.GetSQLite()
	logger = config.GetLogger("Handler")
}
