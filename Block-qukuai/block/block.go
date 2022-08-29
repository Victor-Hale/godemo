package block

import (
	"Block-qukuai/database"
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)
const (
	TABLE_BLOCKS = "blocks"
	BLOCK_LAST   = "last"
)
type Block struct {
	Version       int32  //协议的版本号
	HashPrevBlock []byte //上一个区块的hash值，长度为32个字节
	Time          int32  //时间戳，从1970.01.01 00:00:00到当前时间的秒数
	Bits          int32  //工作量证明(POW)的难度
	Nonce         int32  //要找的符合POW要求的的随机数

	Data          []byte //区块存储的内容，在虚拟币中用来存储交易信息
}

// IntToByte IntToBytes int -> byte将32位整型数字按大端模式转换成字节切片
func IntToByte(val int32) []byte {
	buffer := new(bytes.Buffer)
	binary.Write(buffer, binary.BigEndian, val)
	return buffer.Bytes()
}

// BytesToHex 将字节切片转换成十六进制格式
func BytesToHex(data []byte) string {
	dst := fmt.Sprintf("%x", data)
	return dst
}

// Gethash 创建hash的过程
func (block*Block)Gethash() []byte{
    version := IntToByte(block.Version)
	time := IntToByte(block.Time)
	bits := IntToByte(block.Bits)
	nonce := IntToByte(block.Nonce)
	data := bytes.Join([][]byte{version, block.HashPrevBlock, time, bits, nonce, block.Data}, []byte{})
	hash := sha256.Sum256(data)
	return hash[:]
}

// NewBlock 初始化
func NewBlock(data string, prevHash []byte) *Block {
	block := Block{
		Version:       1.0,
		HashPrevBlock: prevHash,
		Time:          int32(time.Now().Unix()),
		Bits:          0,
		Nonce:         0,
		Data:          []byte(data),
	}
	return &block
}

// NewGenesisBlock 创建创世纪,因为第一个区块没有前区块
func NewGenesisBlock()*Block{
	return NewBlock("Genesis Block", []byte{})
}

//type Object interface{} //节点中存储的数据的类型
////Node 就是一个节点，一个data域，一个Next指向下一个节点的地址
//type Node struct {
//	Data Object //节点存储的数据
//	Next *Node //指向下一个节点的指针
//}
//
//type List struct {
//	HeadNode *Node //头节点
//}
//

//Chain 表示区块链，每一条区块链中包含着多个区块
type Chain struct {
	DB *database.Database
}
type BlockchainIterator struct {
	//当前区块的hash
	hashCurrent []byte
	//区块链数据库
	DB *database.Database
}
// AddBlock 添加一个新的区块
func (bc *Chain) AddBlock(data string) {

	//校验区块链上是否已经有了创世纪区块
	hashLast := bc.DB.Get(TABLE_BLOCKS, BLOCK_LAST)
	if len(hashLast) == 0 {
		//当前链上没有区块直接返回
		return
	}

	//取出当前区块链的最后一个区块
	val := bc.DB.Get(TABLE_BLOCKS, string(hashLast))
	prevBlock := Deserialize(val)

	//传入区块数据和最后一个区块的hash，建新的区块
	block := NewBlock(data, prevBlock.Gethash())

	//取该区块的哈希值
	blockHash := block.Gethash()

	//将该区块的哈希值和序列化数据组成键值对存入数据库
	bc.DB.Set(TABLE_BLOCKS, string(blockHash), block.Serialize())

	//将最后一个区块的哈希值存入数据库，Key 标记为 "last"
	bc.DB.Set(TABLE_BLOCKS, BLOCK_LAST, blockHash)
}

// AddGenesisBlock 向区块链上增加创世纪区块
func (bc *Chain) AddGenesisBlock() {

	//检查区块链上是否已经存在区块
	hashLast := bc.DB.Get(TABLE_BLOCKS, BLOCK_LAST)
	if len(hashLast) > 0 {
		//区块链上已经有区块了，不能添加创世纪区块
		return
	}

	//创建一个创世纪区块
	block := NewGenesisBlock()
	//取该区块的哈希值
	blockHash := block.Gethash()

	//将该区块的哈希值和序列化数据组成键值对存入数据库
	bc.DB.Set(TABLE_BLOCKS, string(blockHash), block.Serialize())

	//将最后一个区块的哈希值存入数据库，Key 标记为 "last"
	bc.DB.Set(TABLE_BLOCKS, BLOCK_LAST, blockHash)
}

// Iterator 获得遍历区块链的迭代子
func (bc *Chain) Iterator() *BlockchainIterator {
	hashLast := bc.DB.Get(TABLE_BLOCKS, BLOCK_LAST)
	if len(hashLast) == 0 {
		//数据库中没有最后区块的记录，说明还没生成区块，无需迭代
		return nil
	}

	//取出最后一个区块的hash，初始化迭代子，准备迭代
	it := &BlockchainIterator{
		hashCurrent:hashLast,
		DB: bc.DB,
	}

	return it
}

// Next 遍历区块链，返回当前区块，移向下一区块
func (it *BlockchainIterator) Next() *Block {
	if it.hashCurrent == nil {
		//已经没有前一区块，到此结束
		return nil
	}
	//取出当前遍历到的区块
	val := it.DB.Get(TABLE_BLOCKS, string(it.hashCurrent))
	block := Deserialize(val)

	it.hashCurrent = block.HashPrevBlock

	return block
}

// GetCount 统计区块链中区块的数量
func (it *BlockchainIterator) GetCount() int {
	var count int = 0
	if it.hashCurrent == nil {
		return count
	}
	for {
		if it.Next() != nil {
			count ++
			continue
		}
		break
	}
	return count
}

// NewBlockChain 新建一个区块链对象
func NewBlockChain() *Chain {
	blockChain := Chain{database.NewDatabase()}
	blockChain.AddGenesisBlock()
	return &blockChain
}

//type Chain1 struct {
//	Blocks []*Block
//}
////新建一个区块链对象
//func NewBlockChain1() *Chain1 {
//	//预先创建一个创世纪区块
//	blockChain := Chain1{Blocks: []*Block{NewGenesisBlock()}}
//
//	return &blockChain
//}

// Serialize 区块序列化，也就是将区块结构的内部数据转换为可以存储的字节流的格式(为了能将block存进数据库，需要对他进行序列化与反序列化)
func (block *Block) Serialize() []byte {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}

// Deserialize 区块反序列化，也就是将字节流转换为含有内部数据的区块结构，这个过程跟Serialize正好相反(为了能将block存进数据库，需要对他进行序列化与反序列化)
func Deserialize(bytesBlock []byte) *Block {
	decoder := gob.NewDecoder(bytes.NewReader(bytesBlock))
	var block Block
	//fmt.Printf("%d\n", decoder)
	err := decoder.Decode(&block)
	if err != nil {
	//	log.Panic(err)
	}

	return &block
}