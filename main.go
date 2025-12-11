package main

import (
	"fmt"
	"log"
)

func main() {
	b := NewBlock("Hello")

	if err := b.Validate(); err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(b)
}