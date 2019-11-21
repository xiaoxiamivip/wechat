package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.Default()

	app.Use(recover.New())
	app.Use(logger.New())
	routeInit(app)
	app.Run(iris.Addr(":8000"))
}
