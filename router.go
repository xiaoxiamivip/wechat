package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/xiaoxiamivip/wechat/database"
	"github.com/xiaoxiamivip/wechat/logic/user"
	"github.com/xiaoxiamivip/wechat/middleware"
)

func routeInit(app *iris.Application) {
	app.Use(database.InitEngine)
	app.Use(middleware.AuthMiddleware)
	users := mvc.New(app.Party("/user"))
	users.Handle(new(user.LoginController))
}
