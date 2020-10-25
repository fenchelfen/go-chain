package gochain

import (
	"gochain/models"
)

type block struct {
	index         int64
	transactions  []*models.Transaction
	prevBlockHash string
}

type blockchain struct {
	blocks []*block
}

func CreateBlock(index int64, transactions []*models.Transaction, prevBlockHash string) *block {

	return &block{index, transactions, prevBlockHash}
}

func CreateBlockChain() *blockchain {

	chain := []*block{
		CreateBlock(0, nil, ""),
	}
	return &blockchain{chain}
}

func (b *block) ComputeHashSum() string {

	return ""
}

func (c *blockchain) LastBlock() *block {

	return c.blocks[len(c.blocks)-1]
}
