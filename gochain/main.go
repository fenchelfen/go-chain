package gochain

import (
	"crypto/sha256"
	"fmt"
	"gochain/models"
	"strconv"
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

	var hash = ""

	for i := 0; i < len(b.transactions); i++ {

		hash = fmt.Sprintf(
			"%x",
			sha256.Sum256([]byte(hash+*b.transactions[i].Content)),
		)
	}

	return hash
}

func (b *block) GetDigest(nonce int) string {

	var hash = ""

	for i := 0; i < len(b.transactions); i++ {

		hash = fmt.Sprintf(
			"%x",
			sha256.Sum256([]byte(hash+*b.transactions[i].Content+strconv.Itoa(nonce))),
		)
	}

	return hash
}

func (c *blockchain) LastBlock() *block {

	return c.blocks[len(c.blocks)-1]
}

func (c *blockchain) ProveWork(b *block) string {

	var nonce = 0

	for ; !c.IsValidProof(b.GetDigest(nonce)); nonce++ {
	}

	return b.GetDigest(nonce)
}

func (c *blockchain) IsValidProof(proof string) bool {

	return strings.HasPrefix(proof, "00")
}

func (c *blockchain) AddBlock(b *block, proof string) bool {

	if c.LastBlock().ComputeHashSum() != b.prevBlockHash {
		return false
	}

	if !c.IsValidProof(proof) {
		return false
	}

	c.blocks = append(c.blocks, b)

	return true
}
