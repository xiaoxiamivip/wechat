package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris/v12"
	"math/rand"
)

var engines []*xorm.Engine

func InitEngine(ctx iris.Context) {
	for i := 0; i < 10; i++ {
		//engine, err := xorm.NewEngine("mysql", "root:123456@/wechat?charset=utf8")
		engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
		if err != nil {
			ctx.Application().Logger().Warnf("database conect err:%s", err)
			continue
		}
		engines = append(engines, engine)
	}
	ctx.Next()
}

func DB() *xorm.Engine {
	if len(engines) == 0 {
		panic("db error")
		return nil
	}
	num := rand.Intn(len(engines))
	return engines[num]
}
