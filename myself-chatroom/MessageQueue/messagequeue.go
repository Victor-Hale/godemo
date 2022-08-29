package Message

import (
	"fmt"
)

type MessageQueue interface {
	send(map[string][]byte)  bool
    get() map[string]byte
}

type Message struct {
	Date map[string][]byte
}

func (M*Message)send (message map[string][]byte) bool{
     M.Date = message
	 return true
}

func (M*Message)get() map[string][]byte{
	for key , value := range M.Date{
		fmt.Println(key,string(value))
	}
	return M.Date
}

//func main(){
//	user := make(map[string][]byte)
//	s:=[]byte("123")
//	user["name"] =s
//    var message Message
//	message.send(user)
//	message.get()
//}