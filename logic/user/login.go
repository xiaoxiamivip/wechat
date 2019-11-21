package user

import (
	"github.com/kataras/iris/v12"
	//"github.com/kataras/iris/v12/mvc"
)

type LoginController struct {
	Ctx iris.Context
}

func (c *LoginController) GetIndexBy(test string) (int, error) {
	return c.Ctx.JSON(iris.Map{"message": test})
}
