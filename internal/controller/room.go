package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nzx.org/music-together/internal/config"
	"nzx.org/music-together/internal/model"
)

type RoomController struct{}

func (rc *RoomController) Register(engine *gin.Engine, db *gorm.DB) {
	v1(engine)
	config.AutoMigrate(&model.Room{})
}

func v1(engine *gin.Engine) {
	v1 := engine.Group("/api/v1/")

	v1.GET("room", listRooms)
	v1.PUT("room", createRoom)
	v1.POST("room", updateRooms)
	v1.DELETE("room", deleteRooms)
}

func listRooms(c *gin.Context) {
	var rooms []model.Room
	tx := config.GetDB().Find(&rooms)
	if tx.Error != nil {
		c.JSON(500, gin.H{"List room failed.": tx.Error.Error()})
	} else {
		c.JSON(200, gin.H{"rooms": rooms})
	}
}

func createRoom(c *gin.Context) {
	var room model.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(400, gin.H{"Create room failed.": err.Error()})
		return
	}
	tx := config.GetDB().Create(&room)
	if tx.Error != nil {
		c.JSON(500, gin.H{"Create room failed.": tx.Error.Error()})
	} else {
		c.JSON(200, gin.H{"room": room})
	}
}

func updateRooms(c *gin.Context) {
	var rooms []model.Room
	if err := c.ShouldBindJSON(&rooms); err != nil {
		c.JSON(400, gin.H{"Update rooms failed.": err.Error()})
		return
	}
	tx := config.GetDB()
	for _, room := range rooms {
		tx = tx.Model(&room).Where("id = ?", room.ID).Updates(room)
		if tx.Error != nil {
			c.JSON(500, gin.H{"Update rooms failed.": tx.Error.Error()})
			return
		}
	}
	c.JSON(200, gin.H{"rooms": rooms})
}

func deleteRooms(c *gin.Context) {
	var ids []uint
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(400, gin.H{"Delete rooms failed.": err.Error()})
		return
	}
	tx := config.GetDB().Delete(&model.Room{}, ids)
	if tx.Error != nil {
		c.JSON(500, gin.H{"Delete rooms failed.": tx.Error.Error()})
	} else {
		c.JSON(200, gin.H{"ids": ids})
	}
}
