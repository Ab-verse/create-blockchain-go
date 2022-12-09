package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce int 
	previousHash string
	timestamp int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block{
	b := new(Block)
	b.nonce = nonce
	b.previousHash = previousHash
	b.timestamp = time.Now().UnixNano()
	return b
}

func (b *Block) Print() {
	fmt.Println("nonse            ", b.nonce)
	fmt.Println("previous_hash    ", b.previousHash)
	fmt.Println("timestamp        ", b.timestamp)
	fmt.Println("transactions     ", b.transactions)	
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonse int, previousHash string) *Block {
	b := NewBlock(nonse, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Blockchain: ")
}
func main() {
	blockChain := NewBlockchain()
	blockChain.Print()
	blockChain.CreateBlock(5, "hash 1")
	blockChain.Print()
	blockChain.CreateBlock(2, "hash 2")
	blockChain.Print()
}