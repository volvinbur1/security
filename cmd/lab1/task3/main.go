package main

import (
	"fmt"
	"github.com/volvinbur1/security/internal/lab1/genetic"
	"os"
	"time"
)

func main() {
	cipher, err := os.ReadFile("./task3.txt")
	if err != nil {
		panic(err)
	}

	genAlg := genetic.New(500, 150, "./trigrams.txt")
	startTime := time.Now()
	plaintext, key := genAlg.Decrypt(cipher)
	fmt.Println("Execution time:", time.Since(startTime).String())
	fmt.Println(string(plaintext))
	fmt.Println("Key:", key)
}
