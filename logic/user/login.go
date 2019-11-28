package user

import (
	"github.com/kataras/iris/v12"
	//"github.com/kataras/iris/v12/mvc"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/xiaoxiamivip/wechat/model/user"
)

type LoginController struct {
	Ctx iris.Context
}

func (c *LoginController) GetIndexBy(test string, id int64) (int, error) {
	user.GetUserInfo()
	err := validate.Struct(mystruct)
	validationErrors := err.(validator.ValidationErrors)
	return c.Ctx.JSON(iris.Map{"message": test, "id": fmt.Sprintf("%d", id)})
}
