package user

import (
	"github.com/kataras/iris/v12"
	//"github.com/kataras/iris/v12/mvc"
	"fmt"
)

type LoginController struct {
	Ctx iris.Context
}

func (c *LoginController) GetIndexBy(test string, id int64) (int, error) {
	return c.Ctx.JSON(iris.Map{"message": test, "id": fmt.Sprintf("%d", id)})
}
