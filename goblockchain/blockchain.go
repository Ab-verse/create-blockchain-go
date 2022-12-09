package main

import (
	"fmt"
	"log"
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


func init() {
	log.SetPrefix("Blockchain: ")
}
func main() {
	b := NewBlock(0, "init hash")
	b.Print()
}