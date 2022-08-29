package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	 resp,error:= http.Get("https://www.docin.com/touch_new/preview_new.do?id=2358625198")   //通过get请求域名
	if error != nil{
    panic(error)  //返回报错信息
	}
	defer resp.Body.Close()  //关闭响应通道
    if resp.StatusCode != http.StatusOK{
    	 fmt.Println("Error:status code err !please careful",resp.StatusCode)//返 回错误信息
		return
    }   //判断访问状态是否正正常
    	 all,error := ioutil.ReadAll(resp.Body)   //读取访问网页所有信息
        	if error != nil{
    		panic(error)    // 打印错误信息
		}
		fmt.Println(string(all))   // 网站信息输出
}
