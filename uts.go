package main

import (
	"crypto/rand"
	"fmt"
)

func GenerateId() string {
	bytes := make([]byte, 16)
	_, _ = rand.Read(bytes)

	return fmt.Sprintf("%x", bytes)
}