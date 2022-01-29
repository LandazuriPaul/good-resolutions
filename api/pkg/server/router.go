package server

import (
	"github.com/LandazuriPaul/good-resolutions/api/internal/orm"
	"github.com/LandazuriPaul/good-resolutions/api/pkg/server/routes"
	"github.com/LandazuriPaul/good-resolutions/api/pkg/utils"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes register the routes for the server
func RegisterRoutes(config *utils.ServerConfig, r *gin.Engine, orm *orm.ORM) (err error) {
	// Auth routes
	if err = routes.Auth(config, r, orm); err != nil {
		return err
	}

	// GraphQL server routes
	if err = routes.GraphQL(config, r, orm); err != nil {
		return err
	}

	// Miscellaneous routes
	if err = routes.Misc(config, r, orm); err != nil {
		return err
	}
	return err
}
