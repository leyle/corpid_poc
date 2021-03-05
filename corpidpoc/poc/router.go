package poc

import (
	"github.com/gin-gonic/gin"
	"github.com/leyle/corpid_poc/corpidpoc/context"
)

func POCRouter(ctx *context.ApiContext, g *gin.RouterGroup) {
	authR := g.Group("", func(c *gin.Context) {
		context.Auth(ctx, c)
	})
	pocR := authR.Group("/poc")
	{
		// credential data schema
		pocR.POST("/credentialdata/new", context.HandlerWrapper(CreateCredentialDataHandler, ctx))
		pocR.GET("/credentialdata/info/:id", context.HandlerWrapper(GetCredentialDataInfoHandler, ctx))
		pocR.GET("/credentialdatas", context.HandlerWrapper(QueryCredentialDataHandler, ctx))

		// authentication did
		pocR.POST("/authenticationdid/new", context.HandlerWrapper(CreateAuthenticationDIDhandler, ctx))
		pocR.GET("/authenticationdid/info/:id", context.HandlerWrapper(GetAuthenticationDIDInfoHandler, ctx))
		pocR.GET("/authenticationdids", context.HandlerWrapper(QueryAuthenticationDIDHandler, ctx))

		// credential did
		pocR.POST("/credentialdid/new", context.HandlerWrapper(CreateCredentialDIDhandler, ctx))
		pocR.GET("/credentialdid/info/:id", context.HandlerWrapper(GetCredentialDIDInfoHandler, ctx))
		pocR.GET("/credentialdids", context.HandlerWrapper(QueryCredentialDIDHandler, ctx))
	}
}
