package gochain

import "gochain/models"

type Blockchain struct {
	index int64
	transactions []*models.Transaction
}

func NewBlockchain() *Blockchain {

	blockchain := new(Blockchain)
	blockchain.index = index
	blockchain.transactions = transactions

	return blockchain;
}