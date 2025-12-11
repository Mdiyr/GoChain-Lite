package main

import (
	"crypto/sha256"
	"fmt"
)

func GenerateHash(data ...any) []byte {
	hasher := sha256.New()

	fmt.Fprint(hasher, data...)

	return hasher.Sum(nil)
}