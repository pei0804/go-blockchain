package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block ブロックチェーン内のブロック
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// NewBlock 新しいブロックを作成する
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock 起点ブロックを作成する
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// SetHash 要素を連結したものをハッシュ化しセットする
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
