package routes

import (
	"github.com/LandazuriPaul/good-resolutions/api/internal/handlers"
	auth "github.com/LandazuriPaul/good-resolutions/api/internal/handlers/auth/middleware"
	"github.com/LandazuriPaul/good-resolutions/api/internal/logger"
	"github.com/LandazuriPaul/good-resolutions/api/internal/orm"
	"github.com/LandazuriPaul/good-resolutions/api/pkg/utils"
	"github.com/gin-gonic/gin"
)

// GraphQL routes
func GraphQL(config *utils.ServerConfig, r *gin.Engine, orm *orm.ORM) error {
	// GraphQL paths
	gqlPath := config.VersionedEndpoint(config.GraphQL.Path)
	pgqlPath := config.GraphQL.PlaygroundPath
	g := r.Group(gqlPath)

	// GraphQL handler
	g.POST("", auth.Middleware(g.BasePath(), config, orm), handlers.GraphqlHandler(orm))
	logger.Info("GraphQL @ ", gqlPath)
	// Playground handler
	if config.GraphQL.IsPlaygroundEnabled {
		logger.Info("GraphQL Playground @ ", g.BasePath()+pgqlPath)
		g.GET(pgqlPath, handlers.PlaygroundHandler(g.BasePath()))
	}

	return nil
}