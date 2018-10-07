package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//定义一个区块的类
type Block struct {
	//时间戳
	TimeStamp int64
	//数据
	data []byte
	//前一个区块的hash
	PreHash []byte
	//这个区块的hash值
	Hash []byte
}

func (block *Block) SetHash() {
	time := []byte(strconv.FormatInt(block.TimeStamp, 10))
	headers := bytes.Join([][]byte{block.PreHash, block.data, time}, []byte{})
	sum256 := sha256.Sum256(headers)
	block.Hash = sum256[:]

}

func NewBlock(data string, preHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), preHash, nil}
	block.SetHash()
	return block
}

//创世区块
func NewGodblock() *Block {
	return NewBlock("zmark", []byte{})
}
