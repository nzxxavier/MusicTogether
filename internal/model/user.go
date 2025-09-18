package model

import (
	"context"

	"gorm.io/gorm"
	"nzx.org/music-together/internal/config"
	"nzx.org/music-together/internal/utils"
)

type User struct {
	gorm.Model `json:"omitempty"`
	UserName   string `json:"userName"`
	UserIp     string `json:"userIp"`
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	id, err := utils.GenerateSnowflakeID()
	if err != nil {
		config.Log().Error(context.Background(), "Generate snowflake id failed:", err)
		return err
	}
	user.ID = id
	return nil
}
