package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp",
		"127.0.0.1:6379",
		redis.DialDatabase(0), //DialOption参数可以配置选择数据库、连接密码等
		redis.DialPassword(""))
	if err != nil {
		fmt.Println("Connect to redis failed ,cause by >>>", err)
		return
	}
	defer conn.Close()
	//set(conn)
	//val := read(conn)
	//fmt.Println(val)
	mymap := make(map[string]string)
	mymap = map[string]string{
		"name" : "xxx",
		"age" : "22",
	}
	for  key,val :=range mymap{
		conn.Do("SET", key ,val)
	}
	result,_:= redis.String(conn.Do("Get","name"))
	fmt.Println(result)
}
func set (conn redis.Conn) string {
	_ , err := conn.Do("SET", "test-Key", "test-Value", "EX", "5")
	if err != nil {
		return "error"
	} else {
		return "sucess"
	}
}
func read (conn redis.Conn) string{
	   val,err:=redis.String(conn.Do("Get","name"))
	if err != nil {
		return "error"
	} else {
		return val
	}
}
