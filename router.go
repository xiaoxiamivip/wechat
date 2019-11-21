package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/xiaoxiamivip/wechat/logic/middleware"
	"github.com/xiaoxiamivip/wechat/logic/user"
)

func routeInit(app *iris.Application) {
	app.Use(middleware.AuthMiddleware)
	users := mvc.New(app.Party("/user"))
	users.Handle(new(user.LoginController))
}
