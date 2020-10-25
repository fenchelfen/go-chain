package gochain

import (
	"crypto/sha256"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"gochain/models"
	"io/ioutil"
	"log"
	"net/http"
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

type Client struct {
	UUID       uuid.UUID
	Peers      []*models.Peer
	Blockchain *blockchain
}

var MyClient Client

func CreateBlock(index int64, transactions []*models.Transaction, prevBlockHash string, proofOfWork string) *block {

	return &block{index, transactions, prevBlockHash, 0, proofOfWork}
}

func CreateBlockchain() *blockchain {

	_uuid := strfmt.UUID(MyClient.UUID.String())
	content := "Empty"
	transaction := models.Transaction{
		Author:  &_uuid,
		Content: &content,
	}

	transactions := []*models.Transaction{&transaction}

	genesisBlock := CreateBlock(0, transactions, "", "")
	genesisBlock.prevBlockHash = genesisBlock.ComputeHashSum()
	genesisBlock.proofOfWork = genesisBlock.ProveWork()

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

func (c *blockchain) IsChainValid() bool {

	var prevHash = c.blocks[0].prevBlockHash

	for _, block := range c.blocks {

		if prevHash != block.prevBlockHash {
			return false
		}

		if IsValidProof(&block.proofOfWork) {
			return false
		}

		prevHash = block.ComputeHashSum()
	}

	return true
}

// Retrieve chains of all peers, check for validity and reach consensus
// by answering this question:
//
//			Whose chain is longer?
//
func (c *blockchain) ReachConsensus() *blockchain {

	var peerChains []*blockchain
	longestChain := c

	for _, peer := range MyClient.Peers {

		res, err := http.Get(*peer.NodeAddress + "/chain")

		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()

		if err != nil {
			log.Fatalln(err)
		}

		peerChain := models.Chain{}
		err = peerChain.UnmarshalBinary(body)

		peerBlockchain := MakeBlockchain(&peerChain)
		peerChains = append(peerChains, peerBlockchain)

		if len(peerBlockchain.blocks) > len(longestChain.blocks) && peerBlockchain.IsChainValid() {

			longestChain = peerBlockchain
		}

		if err != nil {
			log.Fatalln(err)
		}
	}

	return longestChain
}

func IsValidProof(proof *string) bool {

	return strings.HasPrefix(*proof, "00")
}

func MakeBlockchain(c *models.Chain) *blockchain {

	blockchain := &blockchain{}

	for _, cBlock := range c.Chain {

		b := CreateBlock(*cBlock.Index, cBlock.Transactions, cBlock.PrevBlockHash, *cBlock.ProofOfWork)
		blockchain.AddBlock(b)
	}

	return blockchain
}

func (c *blockchain) MakeChain() *models.Chain {
	chain := models.Chain{}
	chain.Peers = MyClient.Peers

	for _, block := range c.blocks {

		cBlock := models.Block{
			Index:         &block.index,
			PrevBlockHash: block.prevBlockHash,
			ProofOfWork:   &block.proofOfWork,
			Transactions:  block.transactions,
		}

		chain.Chain = append(chain.Chain, &cBlock)
	}

	return &chain
}
