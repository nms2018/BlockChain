package blockchain

import (
	"bytes"
	"crypto/sha256"
)

//Block is the structure of each block in the blockchain
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

//BlockChain is the blockchain
type BlockChain struct {
	Blocks []*Block
}

//DeriveHash is responsible to return the hash of combined data
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

//CreateBlock now creates the block with data in string
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

//AddBlock ,after creating blocks, it adds it to the system
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

//Genesis is the block that is responsible to initialize the chain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

//InitBlockChain is used to initialize the BlockChain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
