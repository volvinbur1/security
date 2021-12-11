package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	cipher, err := os.ReadFile("./task3.txt")
	if err != nil {
		panic(err)
	}

	startTime := time.Now()
	//plaintext, key := decrypt.RepeatingKeyXorCipher(cipher, keyLength)
	fmt.Println("Execution time:", time.Since(startTime).String())
	//fmt.Println(string(plaintext))
	//fmt.Println("Key:", key)
}
