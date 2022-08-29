package main

import (
	"Block-qukuai/block"
	"fmt"
)

func main() {
	//新建一条区块链，如果区块链数据库中已经有了数据，将会读取数据库加载进来
	//如果尚未创建数据库，或者数据库为空，那么区块链会自动生成一个创世纪区块
	bc := block.NewBlockChain()
	defer bc.DB.Close()
	//如果区块链中只有1个创世纪区块，我们就再添加3个区块。
	if bc.Iterator().GetCount() == 1 {
		bc.AddBlock("Mini block 01")
		bc.AddBlock("Mini block 02")
		bc.AddBlock("Mini block 03")
		bc.AddBlock("Mini block 04")
	}

	//区块链中应该有4个区块：1个创世纪区块，还有3个添加的区块
	iterator := bc.Iterator()
	for {
		next := iterator.Next()
		if next == nil {
			break
		}
		fmt.Println("前一区块哈希值：", block.BytesToHex(next.HashPrevBlock))
		fmt.Println("当前区块内容为：", string(next.Data))
		fmt.Println("当前区块哈希值：", block.BytesToHex(next.Gethash()))
		fmt.Println("=============================================")
	}

	//bc := block.NewBlockChain()
	//defer bc.DB.Close()
	//bc.DB.Open()
	//bc.DB.Set("blocks","last", []byte("123"))
	//a:=bc.DB.Get("blocks","last")
	//fmt.Printf(string(a))

	//bc:= database.Database{}
	//bc.Open()
	//defer bc.Close()
	//bc.Set("blocks","last", []byte("123"))
	//bc.Set("blocks","last1", []byte("123456"))
	//a:=bc.Get("blocks","last")
	//fmt.Printf(string(a))

}
