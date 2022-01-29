package main

import (
	"github.com/LandazuriPaul/good-resolutions/api/internal/orm"
	"github.com/LandazuriPaul/good-resolutions/api/pkg/logger"
	"github.com/LandazuriPaul/good-resolutions/api/pkg/server"
	"github.com/LandazuriPaul/good-resolutions/api/pkg/utils"
)

func main() {
	var serverConfig = &utils.ServerConfig{
		Host:          utils.MustGetString("SERVER_HOST"),
		Port:          utils.MustGetString("SERVER_PORT"),
		URISchema:     utils.MustGetString("SERVER_URI_SCHEMA"),
		Version:       utils.MustGetString("SERVER_PATH_VERSION"),
		Mode:          utils.MustGetString("SERVER_MODE"),
		SessionSecret: utils.MustGetString("SERVER_SESSION_SECRET"),
		JWT: utils.JWTConfig{
			Secret:    utils.MustGetString("AUTH_JWT_SECRET"),
			Algorithm: utils.MustGetString("AUTH_JWT_SIGNING_ALGORITHM"),
		},
		GraphQL: utils.GraphQLConfig{
			Path:                utils.MustGetString("GRAPHQL_PATH"),
			PlaygroundPath:      utils.MustGetString("GRAPHQL_PLAYGROUND_PATH"),
			IsPlaygroundEnabled: utils.MustGetBool("GRAPHQL_PLAYGROUND_ENABLED"),
		},
		Database: utils.DatabaseConfig{
			Dialect:      utils.MustGetString("GORM_DIALECT"),
			Host:         utils.MustGetString("GORM_HOST"),
			Port:         utils.MustGetString("GORM_PORT"),
			Username:     utils.MustGetString("GORM_USERNAME"),
			Password:     utils.MustGetString("GORM_PASSWORD"),
			DatabaseName: utils.MustGetString("GORM_DATABASE_NAME"),
			SSLMode:      utils.MustGetString("GORM_SSL_MODE"),
			SeedDB:       utils.MustGetBool("GORM_SEED_DB"),
			LogMode:      utils.MustGetBool("GORM_LOGMODE"),
			AutoMigrate:  utils.MustGetBool("GORM_AUTOMIGRATE"),
		},
	}
	orm, err := orm.Factory(serverConfig)
	defer orm.DB.Close()
	if err != nil {
		logger.Panic(err)
	}
	server.Run(serverConfig, orm)
}
