package main

import (
	"fmt"
)

func main() {
           a:="url:\"http://wenzhuhao.oss-cn-beijing.aliyuncs.com/uploadimg%2Fpic-121.jpg?OSSAccessKeyId=LTAI5t8wje5EZwbzKUdFHTVr&Expires=1661527843&Signature=8Y%2BnwQr7pchgqgoMzFTDnGIPCOo%3D\""
		   for k , v :=range a{
			   if v=='?' {
				   fmt.Println(a[5:k])
			   }
		   }



}

