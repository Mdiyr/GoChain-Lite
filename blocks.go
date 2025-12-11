package main

import (
	"bytes"
	"fmt"
	"time"
)

type Block struct {
	Data []byte
	Timestamp time.Time
	Hash []byte
	PrevHash []byte
}

type Blockchain struct {
	Blocks []*Block
}

func (b *Block) String() string {
	return fmt.Sprintf(
		// hex format just for show the hash better not for storage. in storage it should be raw bytes
		"Time: %s\n Data: %s\n Hash: %x\n PrevHash: %x\n",
		b.Timestamp, b.Data, b.Hash, b.PrevHash,
	)
}

func (b *Block) Validate() error {
	h :=GenerateHash(b.Timestamp.UnixNano(), b.Data, b.PrevHash)
	if !bytes.Equal(h, b.Hash) {
		return fmt.Errorf("the hash is invalid for block with data %x", b.Hash)
	}
	return nil
}

func NewBlock(data string, prevHash []byte)	*Block {
	b := Block{
		Timestamp: time.Now(),
		Data:      []byte(data),
		PrevHash:  prevHash,
	}
	b.Hash = GenerateHash(b.Timestamp.UnixNano(), b.Data, b.PrevHash)

	return &b 
}

func NewBlockchain() *Blockchain {
	bc := Blockchain{}
	bc.Blocks = []*Block{NewBlock("Genesis Block", []byte{})}
	return &bc
}

func (bc *Blockchain) AddBlock(data string) {
	if len(bc.Blocks) == 0 {
		panic("something wrong about chain!")
	}
	bc.Blocks = append(bc.Blocks, NewBlock(data, bc.Blocks[len(bc.Blocks)-1].Hash))
}

func (bc *Blockchain) String() string {
	var ret string
	for i := range bc.Blocks{
		ret += bc.Blocks[i].String() + "\n"
	}
	return ret
}

func (bc *Blockchain) Validate() error {
	for i := range bc.Blocks {
		if err := bc.Blocks[i].Validate(); err != nil {
			return fmt.Errorf("Block is not valid! %w", err)
		}
		if i > 0 {
			if !bytes.Equal(bc.Blocks[i].PrevHash, bc.Blocks[i-1].Hash) {
				return fmt.Errorf("the previous hash is invalid for block with data %x", bc.Blocks[i].Data)
			}
		}
	}
	return nil
}