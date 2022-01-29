package routes

import (
	"github.com/LandazuriPaul/good-resolutions/api/internal/handlers/auth"
	"github.com/LandazuriPaul/good-resolutions/api/internal/orm"
	"github.com/LandazuriPaul/good-resolutions/api/pkg/utils"
	"github.com/gin-gonic/gin"
)

// Auth routes
func Auth(cfg *utils.ServerConfig, r *gin.Engine, orm *orm.ORM) error {
	// OAuth handlers
	g := r.Group(cfg.VersionedEndpoint("/auth"))
	g.GET("/:provider", auth.Begin())
	g.GET("/:provider/callback", auth.Callback(cfg, orm))
	// g.GET(:provider/refresh", auth.Refresh(cfg, orm))
	return nil
}
