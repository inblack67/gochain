package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.PrevHash, b.Data}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Data:     []byte(data),
		PrevHash: prevHash,
		Hash:     []byte{},
	}
	block.DeriveHash()
	return block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	block := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, block)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *Blockchain {
	return &Blockchain{
		[]*Block{Genesis()},
	}
}

func main() {

	chain := InitBlockchain()

	chain.AddBlock("First")
	chain.AddBlock("Second")
	chain.AddBlock("Third")

	for _, block := range chain.blocks {
		fmt.Printf("Prev      : %x\n", block.PrevHash)
		fmt.Printf("Data      : %s\n", block.Data)
		fmt.Printf("Hash      : %x\n", block.Hash)
	}

	fmt.Println("hello")

}
