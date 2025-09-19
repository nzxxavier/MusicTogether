package config

import (
	logursLogger "github.com/siruspen/logrus"
	"github.com/spf13/viper"
)

const (
	LogFileName       = "log.fileName"
	LogFileMaxSizeMB  = "log.fileMaxSizeMB"
	LogFileMaxBackups = "log.fileMaxBackups"
	LogFileMaxAgeDays = "log.fileMaxAgeDays"
	LogLevel          = "log.level"
	LogCompress       = "log.compress"

	GinMode     = "gin.mode"
	GinPort     = "gin.port"
	GinAddress  = "gin.address"
	GinLogLevel = "gin.logLevel"

	DataSourceHost               = "dataSource.host"
	DataSourcePort               = "dataSource.port"
	DataSourceUser               = "dataSource.user"
	DataSourcePassword           = "dataSource.password"
	DataSourceDBName             = "dataSource.dbName"
	DataSourceMaxOpenConnections = "dataSource.maxOpenConnections"
)

func InitConfig(configPath *string) {
	initLogDefaultConfigs()
	initGinDefaultConfigs()
	initDataSourceDefaultConfigs()

	viper.SetConfigFile(*configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initLogDefaultConfigs() {
	viper.SetDefault(LogFileName, "music-together.log")
	viper.SetDefault(LogFileMaxSizeMB, 100)
	viper.SetDefault(LogFileMaxBackups, 7)
	viper.SetDefault(LogFileMaxAgeDays, 30)
	viper.SetDefault(LogCompress, true)
}

func initGinDefaultConfigs() {
	viper.SetDefault(GinMode, "release")
	viper.SetDefault(GinPort, 8080)
	viper.SetDefault(GinAddress, "127.0.0.1")
	viper.SetDefault(GinLogLevel, logursLogger.DebugLevel)
}

func initDataSourceDefaultConfigs() {
	viper.SetDefault(DataSourceHost, "127.0.0.1")
	viper.SetDefault(DataSourcePort, 5432)
	viper.SetDefault(DataSourceUser, "music_together")
	viper.SetDefault(DataSourcePassword, "music_together")
	viper.SetDefault(DataSourceDBName, "music_together")
	viper.SetDefault(DataSourceMaxOpenConnections, 10)
}
