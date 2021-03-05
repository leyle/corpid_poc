package context

import (
	"github.com/gin-gonic/gin"
	"github.com/leyle/fabric-user-manager/jwtwrapper"
	"github.com/leyle/fabric-user-manager/model"
	"github.com/leyle/go-api-starter/ginhelper"
	"github.com/rs/zerolog"
)

type ApiContext struct {
	C      *gin.Context
	Cfg    *Config
	JWTCtx *model.JWTContext
}

func (ctx *ApiContext) New(c *gin.Context) *ApiContext {
	jwtCtx := ctx.JWTCtx.New(c)
	n := &ApiContext{
		C:      c,
		Cfg:    ctx.Cfg,
		JWTCtx: jwtCtx,
	}
	return n
}

func (ctx *ApiContext) Logger() *zerolog.Logger {
	logger := zerolog.Ctx(ctx.C.Request.Context())
	return logger
}

func Auth(ctx *ApiContext, c *gin.Context) {
	newCtx := ctx.New(c)
	jwtCtx := newCtx.JWTCtx
	resp := jwtwrapper.Auth(jwtCtx)
	if resp.Err != nil {
		ginhelper.Return401Json(c, resp.Err.Error())
	}
	c.Next()
}

func HandlerWrapper(f func(ctx *ApiContext), ctx *ApiContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		nCtx := ctx.New(c)
		f(nCtx)
	}
}
