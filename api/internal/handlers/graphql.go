package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/LandazuriPaul/good-resolutions/api/internal/gql"
	"github.com/LandazuriPaul/good-resolutions/api/internal/gql/resolvers"
	"github.com/LandazuriPaul/good-resolutions/api/internal/orm"
	"github.com/gin-gonic/gin"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler(orm *orm.ORM) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	c := gql.Config{
		Resolvers: &resolvers.Resolver{ORM: orm},
	}

	h := handler.GraphQL(gql.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler Defines the Playground handler to expose our playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := handler.Playground("Good Resolutions API", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
