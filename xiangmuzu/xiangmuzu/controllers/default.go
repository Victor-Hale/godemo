package controllers

import (
	"github.com/astaxie/beego"
	"xiangmuzu/xiangmuzu/global"
	"xiangmuzu/xiangmuzu/jwt"
	"xiangmuzu/xiangmuzu/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type MySqlController struct {
	JsonController
}

// Registered 添加 注册账户  has256加密
func (c *MySqlController) Registered() {
	Username :=c.GetString("username")
	Password :=c.GetString("password")
	abs := models.Jwt{}
	//srcByte := []byte(Password)
	//sha256New := sha256.New()
	//sha256Bytes := sha256New.Sum(srcByte)
	//Password = hex.EncodeToString(sha256Bytes)
	Password, err :=jwt.HashAndSalt(Password)
	usernamecount :=models.Username{}
	var count int64
	global.Db.Model(usernamecount).Where("username",Username).Count(&count)
	if count==0{
		if err!=nil{
			panic(err)
		}else {
			result := global.Db.Model(abs).Create(map[string]interface{}{
				"username": Username,
				"password": Password,
			})
			if result!=nil{
				c.ApiJsonReturn("Registered successfully",200,"ok")
			} else{
				c.ApiJsonReturn("Registration failed",100,false)
			}
		}
	}else{
		c.ApiJsonReturn("he account has been registered. Please apply for a new account",100,false)
	}
}
type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
}
// Login 获取token(登陆)
func (c *MySqlController) Login(){
	Username :=c.GetString("username")
	Passwordget :=c.GetString("password")
	//srcByte := []byte(Passwordget)
	//sha256New := sha256.New()
	//sha256Bytes := sha256New.Sum(srcByte)
	//Passwordget = hex.EncodeToString(sha256Bytes)

	Password := models.Password{}
	global.Db.Where("username",Username).First(&Password)
	Passwordsql := Password.Password  //数据库密码加密后的hash256加盐
	user := User{Username, Passwordget}
	token,err := jwt.GenerateToken((*jwt.User)(&user),0) //获取token
	//date,_:=jwt.ValidateToken(token)  //解密出token的password
	date :=jwt.ComparePasswords(Passwordsql,Passwordget)
	if date ==false{
		c.ApiJsonReturn("login fail！please check your password or username！",100,false)
	} else{
		if err!=nil{
			panic(err)
		}else {
			c.ApiJsonReturn("login success",200,token)
		}
	}
}

// Get 查询
func (c *MySqlController) Get() {
	user := []models.Information{}
	// user := []utils.User{}
	global.Db.Find(&user)
	if user!=nil{
		c.ApiJsonReturn("请求成功",200,user)
	} else{
		c.ApiJsonReturn("请求失败",100,false)
	}
}

// Liege  模糊查询
func (c *MySqlController) Liege() {
	selecallnode :=c.GetString("selecallnode")
	Title := []models.Information{}
	// user := []utils.User{}
	global.Db.Where("title LIKE ?", selecallnode+"%").Or("anthor LIKE ?", selecallnode+"%").Find(&Title)
	if Title!=nil{
		c.ApiJsonReturn("请求成功",200,Title)
	} else{
		c.ApiJsonReturn("请求失败",100,false)
	}
}