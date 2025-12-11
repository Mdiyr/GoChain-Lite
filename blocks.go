package main

import (
	"fmt"
	"time"
)

type Block struct {
	Data []byte
	Timestamp time.Time
}

func (b *Block) PrintData() string {
	return fmt.Sprintf(
		"Time: %s\n Data: %s\n",
		b.Timestamp, b.Data,
	)
}

func NewBlock(data string)	*Block {
	b := Block{
		Timestamp: time.Now(),
		Data:      []byte(data),
	}
	return &b 
}