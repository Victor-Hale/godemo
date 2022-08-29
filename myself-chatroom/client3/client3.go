package main
import (
	"fmt"
	"myself-chatroom/MessageQueue"
)

func main() {
   DB :=Message.NewDatabase()
   DB.Open()
	defer DB.Close()
   DB.Set("blocks","last", []byte("123"))
   DB.Set("blocks","last1", []byte("1234"))
   DB.Set("blocks","last2", []byte("12345"))
   value := DB.Gatall()
   fmt.Println(value)
}
