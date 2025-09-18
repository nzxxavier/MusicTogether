package config

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", viper.GetString(DataSourceHost),
		viper.GetInt(DataSourcePort), viper.GetString(DataSourceUser), viper.GetString(DataSourcePassword),
		viper.GetString(DataSourceDBName))
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   Log(),
	})
	if err != nil {
		Log().Error(context.Background(), "Failed to connect to database:", err)
	}
}

func AutoMigrate(models ...interface{}) {
	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			Log().Error(context.Background(), "Failed to auto migrate:", err)
		}
	}
}

func GetDB() *gorm.DB {
	return db
}
