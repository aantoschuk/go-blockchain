package main

// It's an ordered back-linked list
// TODO: implement map to keep hash -> block pairs.
type BlockChain struct {
	blocks []*Block
}

func (bc *BlockChain) AddBlock(data string) {
	// get latest block from array
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	// add new element to the end
	bc.blocks = append(bc.blocks, newBlock)
}

// To add a new Block there is a need to have already existing Block
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}
