package gochain

import (
	"gochain/models"
)

type block struct {
	index         int64
	transactions  []*models.Transaction
	prevBlockHash string
}

func CreateBlock(index int64, transactions []*models.Transaction, prevBlockHash string) *block {

	return &block{index, transactions, prevBlockHash}
}

func (b *block) computeHashSum() string {

	return ""
}
