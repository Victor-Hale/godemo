package main

import (
	"fmt"
	"github.com/go-redis/redis"
)


func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379", // url
		Password: "",
		DB:0,   // 0号redis数据库
	})
	result, err := client.Ping().Result()
	if err != nil {
		fmt.Println("ping err :",err)
		return
	}
	fmt.Println(result)
	defer client.Close()
	//set(client)
	// value :=get(client)
	// fmt.Println(value)
	//  result0 := deleted(client)
	//  fmt.Println(result0)
}

func set(Client *redis.Client) *redis.StatusCmd {
     result:=Client.Set("key","val",0)
	 return result    //set key val: OK
}
func get(Client *redis.Client) string {
	value, err := Client.Get("key").Result()
	if err!=nil {
		return "err"
	}
	return value
}
func deleted(Client *redis.Client) string{
		Client.Del("key", "val")
		return "删除成功"
}