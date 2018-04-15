package main

// Blockchain 順序付けされたバックリンクリスト
type Blockchain struct {
	blocks []*Block
}

// NewBlockchain ブロックチェーンを作成する
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// AddBlock ブロックを追加する
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
