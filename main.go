package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
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

func calculateHashForBlock(block Block) string {
	return calculateHash(block.Index, block.PreviousHash, block.timestamp, block.data)
}

func calculateHash(nextIndex int, previousHash string, timestamp time.Time, data string) string {
	str := string(nextIndex) + previousHash + string(timestamp.Unix()) + data
	b := sha256.Sum256([]byte(str))
	return hex.EncodeToString(b[:])
}

func isValidNewBlock(newBlock Block, previousBlock Block) error {
	if previousBlock.Index+1 != newBlock.Index {
		return errors.New("invalid index")
	} else if previousBlock.hash != newBlock.PreviousHash {
		return errors.New("invalid previousHash")
	} else if calculateHashForBlock(newBlock) != newBlock.hash {
		return fmt.Errorf("invalid hash %v %v", calculateHashForBlock(previousBlock), newBlock.hash)
	}
	return nil
}

func isValidChain(blockchainToValidate []Block) error {
	if len(blockchainToValidate) == 0 {
		return errors.New("empty")
	}

	if blockchainToValidate[0].hash != getGenesisBlock().hash {
		return errors.New("invalid genesisBlock")
	}

	previousBlock := blockchainToValidate[0]
	for i := 1; i < len(blockchainToValidate); i++ {
		err := isValidNewBlock(blockchainToValidate[i], previousBlock)
		if err != nil {
			return err
		}
		previousBlock = blockchainToValidate[i]
	}
	return nil
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

	err := isValidChain(blockChain)
	if err == nil {
		fmt.Println("chain is valid")
	} else {
		fmt.Printf("chain error %v\n", err)
	}
}
