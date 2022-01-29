package server

import (
	"github.com/gin-gonic/gin"

	"github.com/LandazuriPaul/good-resolutions/api/internal/handlers"
	"github.com/LandazuriPaul/good-resolutions/api/internal/orm"
	"github.com/LandazuriPaul/good-resolutions/api/pkg/logger"
	"github.com/LandazuriPaul/good-resolutions/api/pkg/utils"
)

// Run spins up the server
func Run(config *utils.ServerConfig, orm *orm.ORM) {
	endpoint := "http://" + host + ":" + port

	r := gin.Default()
	// Handlers
	// Simple keep-alive/ping handler
	r.GET("/ping", handlers.Ping())

	// GraphQL handlers
	// Playground handler
	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		logger.Info("GraphQL Playground @ " + endpoint + gqlPgPath)
	}
	// Pass in the ORM instance to the GraphqlHandler
	r.POST(gqlPath, handlers.GraphqlHandler(orm))
	logger.Info("GraphQL @ " + endpoint + gqlPath)

	// Run the server
	// Inform the user where the server is listening
	logger.Info("Running @ " + endpoint)
	// Print out and exit(1) to the OS if the server cannot run
	logger.Fatal(r.Run(host + ":" + port))
}
