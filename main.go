package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index        int
	PreviousHash string
	timestamp    time.Time
	data         string
	hash         string
}

var blockChain []Block

func getGenesisBlock() Block {
	c := Block{0, "0", time.Unix(1465154705, 0), "my genesis block!!", "816534932c2b7154836da6afc367695e6337db8a921823784c14378abed4f7d7"}
	return c
}

func getBlocks() []Block {
	return blockChain
}

func mineBlock() {

}

func generateNextBlock(data string) Block {
	previuosBlock := getLatestBlock()
	nextIndex := previuosBlock.Index + 1
	nextTimeStamp := time.Now()
	nextHash := calculateHash(nextIndex, previuosBlock.hash, nextTimeStamp, data)
	return Block{nextIndex, previuosBlock.hash, nextTimeStamp, data, nextHash}
}

func getLatestBlock() Block {
	return blockChain[len(blockChain)-1]
}

func calculateHash(nextIndex int, previousHash string, timestamp time.Time, data string) string {
	str := string(nextIndex) + previousHash + string(timestamp.Unix()) + data
	b := sha256.Sum256([]byte(str))
	return hex.EncodeToString(b[:])
}

func main() {
	blockChain = []Block{}
	blockChain = append(blockChain, getGenesisBlock())

	c1 := generateNextBlock("aaaaaa")
	blockChain = append(blockChain, c1)

	c2 := generateNextBlock("bbbbbbbb")
	blockChain = append(blockChain, c2)

	for i := 0; i < len(blockChain); i++ {
		fmt.Printf("%v\n", blockChain[i])
	}

}
