package config

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var ginEngine *gin.Engine

func InitGinEngine() {
	gin.SetMode(viper.GetString(GinMode))
	ginEngine = gin.New()
	ginEngine.Use(GinLogger())
}

func StartGinEngine() {
	if err := ginEngine.Run(viper.GetString(GinAddress) + ":" + viper.GetString(GinPort)); err != nil {
		Log().Error(context.Background(), "Failed to start Gin engine:", err)
	}
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetGin() *gin.Engine {
	return ginEngine
}
