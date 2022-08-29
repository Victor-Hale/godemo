package controllers

import "github.com/astaxie/beego"


////先封装一个json返回
//定义一个Base控制器的结构体
type JsonController struct {
	beego.Controller
}

// JsonReturn 定义一个json返回的结构体
type JsonReturn struct {
	Msg  string 	    `json:"msg"`  //打tag标签
	Code int		    `json:"code"`
	Data interface{}	`json:"data"`
	//Data字段需要设置为interface类型以便接收任意数据
	//json标签意义是定义此结构体解析为json或序列化输出json时value字段对应的key值,如不想此字段被解析可将标签设为`json:"-"`
}

//调用basecontroller的控制器与json的结构体
func (c *JsonController) ApiJsonReturn(msg string,code int,data interface{}) {
	var JsonReturn JsonReturn
	JsonReturn.Msg = msg
	JsonReturn.Code = code
	JsonReturn.Data = data
	c.Data["json"] = JsonReturn		//将结构体数组根据tag解析为json
	c.ServeJSON()					//对json进行序列化输出
	c.StopRun() //终止执行逻辑
}