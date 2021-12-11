package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sourceCodes, err := os.ReadFile("./lab1_cipher.txt")
	if err != nil {
		panic(err)
	}

	encodedStr, err := convertBytesToChars(sourceCodes)
	if err != nil {
		panic(err)
	}

	decodedStr, _ := base64.StdEncoding.DecodeString(encodedStr)

	//err = os.WriteFile("lab1_tasks.txt", decodedStr, 0666)
	//if err != nil {
	//	panic(err)
	//}
	fmt.Println(string(decodedStr))

}

func convertBytesToChars(input []byte) (string, error) {
	chByte := byte(0)
	output := ""
	for i := 0; i < len(input); i++ {
		bit, err := strconv.Atoi(string(input[i]))
		if err != nil {
			return "", err
		}

		if i%8 == 0 && i != 0 {
			output += string(chByte)
			chByte = 0
		}

		chByte <<= 1
		chByte |= byte(bit)
		//fmt.Printf("%b\n", chByte)
	}

	return output, nil
}
