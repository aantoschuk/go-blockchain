package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Timestamp - when the block is created
// Data - is actuall valuable information containing in the block
// PrevBlockHash - stores the hash of the previous block
// Hash - block headers
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// TODO: Make adding new block difficult
// to prevent modification after they're added
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// Function to create new block
// which simplifies the creation of a block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// Genesis Block is a first block in the blockchain and
// is almost always hardcoded
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
