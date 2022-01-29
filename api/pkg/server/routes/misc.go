package routes

import (
	"github.com/LandazuriPaul/good-resolutions/api/internal/handlers"
	"github.com/LandazuriPaul/good-resolutions/api/internal/handlers/auth/middleware"
	"github.com/LandazuriPaul/good-resolutions/api/internal/orm"
	"github.com/LandazuriPaul/good-resolutions/api/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Misc routes
func Misc(config *utils.ServerConfig, r *gin.Engine, orm *orm.ORM) error {
	// Simple keep-alive/ping handler
	r.GET(config.VersionedEndpoint("/ping"), handlers.Ping())
	r.GET(config.VersionedEndpoint("/secure-ping"),
		middleware.Middleware(config.VersionedEndpoint("/secure-ping"), config, orm), handlers.Ping())
	return nil
}