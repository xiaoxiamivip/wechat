package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/go-xorm/xorm"
	"github.com/xiaoxiamivip/wechat/database"
)

type User struct {
	ID        int64  // auto-increment by-default by xorm
	Version   string `xorm:"varchar(200)"`
	Password  string `xorm:"varchar(200)"`
	Languages string `xorm:"varchar(200)"`
}

func GetUserInfoTest() {

	user := User{}
	ok, err := database.DB().Get(&user)

	fmt.Printf("ok:%v user%v err %s", ok, user, err)
}
