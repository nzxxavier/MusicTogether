package model

import (
	"context"

	"gorm.io/gorm"
	"nzx.org/music-together/internal/config"
	"nzx.org/music-together/internal/utils"
)

type Room struct {
	gorm.Model    `json:"omitempty"`
	Password      string `json:"password"`
	AdminPassword string `json:"adminPassword"`
	ExpiredAt     int64  `json:"expiredAt"`
}

func (room *Room) BeforeCreate(db *gorm.DB) error {
	id, err := utils.GenerateSnowflakeID()
	if err != nil {
		config.Log().Error(context.Background(), "Generate snowflake id failed:", err)
		return err
	}
	room.ID = id
	return nil
}
