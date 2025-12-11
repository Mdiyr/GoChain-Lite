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
}

func (b *Block) String() string {
	return fmt.Sprintf(
		// hex format just for show the hash better not for storage. in storage it should be raw bytes
		"Time: %s\n Data: %s\n Hash: %x\n",
		b.Timestamp, b.Data, b.Hash,
	)
}

func (b *Block) Validate() error {
	h :=GenerateHash(b.Timestamp.UnixNano(), b.Data)
	if !bytes.Equal(h, b.Hash) {
		return fmt.Errorf("the hash is invalid for block with data %x", b.Hash)
	}
	return nil
}

func NewBlock(data string)	*Block {
	b := Block{
		Timestamp: time.Now(),
		Data:      []byte(data),
	}
	b.Hash = GenerateHash(b.Timestamp.UnixNano(), b.Data)

	return &b 
}