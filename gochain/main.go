package gochain

import (
	"crypto/sha256"
	"fmt"
	"gochain/models"
	"strings"
)

type block struct {
	index         int64
	transactions  []*models.Transaction
	prevBlockHash string
	nonce         int64
}

type blockchain struct {
	blocks []*block
}

func CreateBlock(index int64, transactions []*models.Transaction, prevBlockHash string) *block {

	return &block{index, transactions, prevBlockHash, 0}
}

func CreateBlockChain() *blockchain {

	genesisBlock := CreateBlock(0, nil, "")
	genesisBlock.prevBlockHash = genesisBlock.ComputeHashSum()

	chain := []*block{
		genesisBlock,
	}
	return &blockchain{chain}
}

func (b *block) ComputeHashSum() string {

	return ""
}

func (b *block) getDigest(nonce int) string {

	var hash = ""

	for i := 0; i < len(b.transactions); i++ {

		hash = fmt.Sprintf(
			"%x",
			sha256.Sum256([]byte(hash)),
		)
	}

	return hash
}

func (c *blockchain) LastBlock() *block {

	return c.blocks[len(c.blocks)-1]
}

func (c *blockchain) ProveWork(b *block) string {

	var nonce = 0

	for strings.HasPrefix(b.getDigest(nonce), "00") {

		nonce++
	}

	return ""
}
