package main

import (
	"encoding/hex"
	"fmt"
	"github.com/volvinbur1/security/internal/lab1/decrypt"
	"os"
	"time"
)

func main() {
	cipherHex, err := os.ReadFile("./lab1_task1.txt")
	if err != nil {
		panic(err)
	}

	cipher, err := hex.DecodeString(string(cipherHex))
	if err != nil {
		panic(err)
	}

	startTime := time.Now()
	plainText, key := decrypt.BruteForce(cipher)
	fmt.Println("Brute force time:", time.Since(startTime).String())
	fmt.Println(string(plainText))
	fmt.Println("Encryption key:", key)
}
