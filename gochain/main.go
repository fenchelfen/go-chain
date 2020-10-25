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
	proofOfWork   string
}

type blockchain struct {
	blocks []*block
}

func CreateBlock(index int64, transactions []*models.Transaction, prevBlockHash string, proofOfWork string) *block {

	return &block{index, transactions, prevBlockHash, 0, proofOfWork}
}

func CreateBlockchain() *blockchain {

	genesisBlock := CreateBlock(0, nil, "", "")
	genesisBlock.prevBlockHash = genesisBlock.ComputeHashSum()
	genesisBlock.proofOfWork = genesisBlock.ProveWork()

	chain := []*block{
		genesisBlock,
	}
	return &blockchain{chain}
}

type proxyObjChain models.Chain

func (c *proxyObjChain) MakeBlockchain() *blockchain {

	blockchain := &blockchain{}

	for _, cBlock := range c.Chain {

		b := CreateBlock(*cBlock.Index, cBlock.Transactions, cBlock.PrevBlockHash, *cBlock.ProofOfWork)
		blockchain.AddBlock(b)
	}

	return blockchain
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

func (b *block) GetDigest(nonce int) *string {

	var hash = ""

	for i := 0; i < len(b.transactions); i++ {

		hash = fmt.Sprintf(
			"%x",
			sha256.Sum256([]byte(hash+*b.transactions[i].Content+strconv.Itoa(nonce))),
		)
	}

	return &hash
}

func (b *block) ProveWork() string {

	var nonce = 0

	for ; !IsValidProof(b.GetDigest(nonce)); nonce++ {
	}

	return *b.GetDigest(nonce)
}

func (c *blockchain) LastBlock() *block {

	return c.blocks[len(c.blocks)-1]
}

func (c *blockchain) AddBlock(b *block) bool {

	if c.LastBlock().ComputeHashSum() != b.prevBlockHash {
		return false
	}

	if !IsValidProof(&b.proofOfWork) {
		return false
	}

	c.blocks = append(c.blocks, b)

	return true
}

func (c *blockchain) CheckChainValidity() bool {

	var prevHash = ""

	for _, block := range c.blocks {

		if prevHash != block.prevBlockHash {
			return false
		}

		if IsValidProof(&block.proofOfWork) {
			return false
		}

		prevHash = block.ComputeHashSum()
	}

	return false
}

func IsValidProof(proof *string) bool {

	return strings.HasPrefix(*proof, "00")
}
