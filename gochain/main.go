package gochain

import (
	"gochain/models"
)

type Blockchain struct {
	index        int64
	transactions []*models.Transaction
	hashSum      string
}

func NewBlockchain(index int64, transactions []*models.Transaction) *Blockchain {

	b := &Blockchain{index, transactions, ""}
	b.computeHashSum()

	return b
}

func (b *Blockchain) computeHashSum() {

	b.hashsum = ""
}
