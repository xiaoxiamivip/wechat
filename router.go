package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/xiaoxiamivip/wechat/logic/user"
)

func routeInit(app *iris.Application) {
	mvc.New(app.Party("/user")).Handle(new(user.LoginController))

}
