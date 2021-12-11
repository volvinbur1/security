package main

import (
	"encoding/base64"
	"fmt"
	"github.com/volvinbur1/security/internal/lab1/decrypt"
	"os"
	"time"
)

const keyLength = 3 // by frequency counting

func main() {
	cipherEncoded, err := os.ReadFile("./lab1_task2.txt")
	if err != nil {
		panic(err)
	}

	cipher, err := base64.StdEncoding.DecodeString(string(cipherEncoded))
	if err != nil {
		panic(err)
	}

	// guess key length
	//frequencies := decrypt.FrequencyAnalyzer(cipher)
	//fmt.Println(frequencies)

	startTime := time.Now()
	plaintext := decrypt.RepeatingKeyXorCipher(cipher, keyLength)
	fmt.Println("Execution time:", time.Since(startTime).String())
	fmt.Println(string(plaintext))
}
