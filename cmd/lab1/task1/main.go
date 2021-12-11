package main

import (
	"encoding/hex"
	"fmt"
	"github.com/volvinbur1/security/internal/lab1/decrypt"
	"os"
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

	possibleKeyValues := ""
	for i := 48; i <= 122; i++ {
		if (i >= 58 && i <= 64) ||
			(i >= 91 && i <= 96) {
			continue
		}
		possibleKeyValues += string(byte(i))
	}

	plainText := decrypt.BruteForce(cipher, possibleKeyValues)
	fmt.Println(string(plainText))
}
