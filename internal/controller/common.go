package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller interface {
	Register(engine *gin.Engine, db *gorm.DB)
}
