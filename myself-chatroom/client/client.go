package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"myself-chatroom/log"
	"net"
)

const (
	HOST   = "localhost"
	PORT   = "8000"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	Message  string `json:"message"`
}
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")
    //1.连接
	 //2.处理与业务逻辑(数据的传输与接收)
	 //3.关闭连接
	conn , err := net.Dial("tcp",HOST+":"+PORT)
	if err !=nil {
		log.Error.Println("连接失败；",err)
		return
	}
	defer conn.Close()
	//conn.RemoteAddr()  获取连接地址
	fmt.Println("您已经成功连接",conn.RemoteAddr())
	var user User
	//fmt.Println("请输入您的个人信息--用户昵称")
	//fmt.Scanln(&user.Name)
	//fmt.Println("请输入您的个人信息--年龄")
	//fmt.Scanln(&user.Age)
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("input name")
	nameEntry.OnChanged = func(content string) {
		fmt.Println("name:", nameEntry.Text, "entered")
	}

	ageEntry := widget.NewEntry()
	ageEntry.SetPlaceHolder("input age")

	nameBox := container.NewHBox(widget.NewLabel("Name"), layout.NewSpacer(), nameEntry)
	ageBox := container.NewHBox(widget.NewLabel("Age"), layout.NewSpacer(), ageEntry)
	user.Name = nameEntry.Text
	user.Age = ageEntry.Text
		for  {
			go func() {
				for {
					Data := make([]byte, 255)
					n,err := conn.Read(Data)
					if n == 0 || err != nil {
						// 收到的参数错误忽略、
						continue
					}
					fmt.Println(string(Data[0:n]))
				}
			}()
			Date := make([]byte,255)
			textEntry := widget.NewEntry()
			textEntry.SetPlaceHolder("input age")
			textBox := container.NewHBox(widget.NewLabel("text"), layout.NewSpacer(), textEntry)
			loginBtn := widget.NewButton("Enter", func() {
				fmt.Println("name:", nameEntry.Text, "age:", ageEntry.Text, "输入了",textEntry.Text)
				Date = []byte(nameEntry.Text + "|" + ageEntry.Text + "|" + textEntry.Text)
				_ , err = conn.Write(Date)

			})
		if err != nil {
			log.Error.Println("...........客户端发送失败；",err)
		}
		fmt.Println("............发送成功")
			content := container.NewVBox(nameBox, ageBox,textBox,loginBtn)
			myWindow.SetContent(content)
			myWindow.Resize(fyne.NewSize(150, 150))
			myWindow.ShowAndRun()
			continue

	   }





}

//
//package main
//
//import (
//	"bufio"
//	"fmt"
//	"net"
//	"os"
//)
//
//var msg = make(chan string)
//
//func read() {
//	for {		//从msg里读出信息并将信息写到控制台
//		m := <-msg
//		fmt.Print(m)
//	}
//}
//
//func write(conn net.Conn) {
//	for {		//读取server端发来地信息,并将信息写到全局的msg里
//		s := make([]byte, 10, 10)
//
//		n, err := conn.Read(s)
//		if err != nil {
//			fmt.Println("err:",err)
//		}
//
//		msg <- string(s[:n])
//
//	}
//
//}
//
//func main() {
//
//	conn, err := net.Dial("tcp", "127.0.0.1:8000")		//发送连接请求
//	if err != nil {
//		fmt.Println("err: ",err)
//	}
//	defer conn.Close()
//	go read()
//	go write(conn)
//	for {		//读取用户从控制台输入地信息
//		reader := bufio.NewReader(os.Stdin)
//		s, err := reader.ReadString('\n')
//		if err != nil {
//			fmt.Println("err:",err)
//
//		}
//		if s == "exit\r\n" {  //用户退出
//
//			conn.Write([]byte(s))
//
//			break
//		} else if s == "change\r\n" {		//用户更名
//
//			fmt.Println("请输入要更改的名字")
//			s, _ := reader.ReadString('\n')
//			s2 := []byte(s[0 : len(s)-2])
//			s3 := append([]byte("change "), s2...)
//			conn.Write([]byte(s3))
//			continue
//
//		}
//		conn.Write([]byte(s))		//将消息发送给server端
//	}
//
//}