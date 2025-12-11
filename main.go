package main

import (
	"fmt"
	"log"
)

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Mhdyr")
	bc.AddBlock("Mrd")
	
	if err := bc.Validate(); err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(bc)
}