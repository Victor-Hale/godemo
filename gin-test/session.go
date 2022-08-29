package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
)

// 初始化一个cookie存储对象
// session-secret是密匙
var store = sessions.NewCookieStore([]byte("session-secret"))

func main() {
	http.HandleFunc("/save", SaveSession)
	http.HandleFunc("/get", GetSession)
	http.HandleFunc("/delete", RemoveSession)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}

// 写入 session
func SaveSession(w http.ResponseWriter, r *http.Request) {
	//　获取一个session对象，session-name是session的名字
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 在session中存储值
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	// 保存更改
	session.Save(r, w)
}

// 读取 session
func GetSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	foo := session.Values["foo"]
	fmt.Println(foo)
}

// 删除 session
func RemoveSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置session的最大存储时间小于零，即删除
	session.Options.MaxAge = -1
	session.Save(r, w)
}