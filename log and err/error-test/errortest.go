package main

import (
	"errors"
	"fmt"
)

func errtest(a int)(string,error){
	if a<=6 {
		return "100",errors.New("not good to find ")
	}else {
		return  "200",nil
	}


}

func main() {
    a:=4
	res, err := errtest(a)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(res)
	}
	fmt.Println("111111")
}
