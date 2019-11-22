package user

import (
	"fmt"
	"github.com/xiaoxiamivip/wechat/database"
)

type User struct {
	ID        int64  // auto-increment by-default by xorm
	Version   string `xorm:"varchar(200)"`
	Salt      string
	Username  string
	Password  string `xorm:"varchar(200)"`
	Languages string `xorm:"varchar(200)"`
}

func GetUserInfo() {
	user := User{}
	ok, err := database.DB().Get(&user)

	fmt.Printf("ok:%v err %s", ok, err)
}
