package main

import (
	"errors"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"xiangmuzu/xiangmuzu/global"
	_ "xiangmuzu/xiangmuzu/routers"
)

var err error
func init(){

	dsn := "xiangmuzu:123456@tcp(139.196.106.241)/xiangmuzu?charset=utf8mb4&parseTime=True&loc=Local"
	global.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		errors.New("err，数据库连接失败")
	}

}
func main() {
	beego.BConfig.WebConfig.AutoRender = false
	beego.Run()
}

