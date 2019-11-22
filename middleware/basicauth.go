package middleware

import (
	"github.com/kataras/iris/v12"
)

func AuthMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
