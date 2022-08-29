package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"net/http"
	"strings"
	"xiangmuzu/xiangmuzu/controllers"
	"xiangmuzu/xiangmuzu/jwt"
)

func init() {
	/*
	  配置跨域
	*/
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		//AllowOrigins:      []string{"https://192.168.0.102"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	var ApiAuthFilter = func(ctx *context.Context) {
		token := ctx.Request.Header.Get("Authorization")
		kv := strings.Split(token, " ")
		if len(kv) != 2 || kv[0] != "Bearer" {
		//	http.Error(ctx.ResponseWriter, "Token verification not pass", http.StatusUnauthorized)
			panic("AuthString无效")
		}
		tokenString := kv[1]
		//fmt.Printf(tokenString)
		_,err:=jwt.ValidateToken(tokenString)
		//fmt.Println(err)
		if err!=nil {
			http.Error(ctx.ResponseWriter, "Token verification not pass", http.StatusUnauthorized)
		} else {
			return
		}
	}


    beego.Router("/", &controllers.MainController{})
	beego.InsertFilter("/token/*",beego.BeforeRouter,ApiAuthFilter)

	beego.Router("/registered", &controllers.MySqlController{},"post:Registered")

	login:= beego.NewNamespace("/api",
		    beego.NSNamespace("/user",
			beego.NSRouter("/registered",&controllers.MySqlController{},"post:Registered"),
			beego.NSRouter("/login",&controllers.MySqlController{},"post:Login")))
	test:=  beego.NewNamespace("/token",
		    beego.NSNamespace("/api",
			beego.NSRouter("/get",&controllers.MySqlController{},"post:Get"),
			beego.NSRouter("/liege",&controllers.MySqlController{},"post:Liege")))

	beego.AddNamespace(test)
	beego.AddNamespace(login)
}

