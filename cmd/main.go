package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nzx.org/music-together/internal/config"
	"nzx.org/music-together/internal/controller"
)

func main() {
	flag.Parse()

	configFile := flag.String("config", "./configs/config.yaml", "Path to config file")
	config.InitConfig(configFile)
	config.InitLogger()
	config.InitDatabase()
	config.InitGinEngine()

	initController(config.GetGin(), config.GetDB())

	config.StartGinEngine()
}

func initController(gin *gin.Engine, db *gorm.DB) {
	var controllers []controller.Controller
	controllers = append(controllers, &controller.RoomController{})

	for _, controllerImplements := range controllers {
		controllerImplements.Register(gin, db)
	}
}
