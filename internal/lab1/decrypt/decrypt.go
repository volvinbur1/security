package decrypt

var charsScore = map[byte]int{
	byte('U'): 2, byte('u'): 2,
	byte('L'): 3, byte('l'): 3,
	byte('D'): 4, byte('d'): 4,
	byte('R'): 5, byte('r'): 5,
	byte('H'): 6, byte('h'): 6,
	byte('S'): 7, byte('s'): 7,
	byte(' '): 8,
	byte('N'): 9, byte('n'): 9,
	byte('I'): 10, byte('i'): 10,
	byte('O'): 11, byte('o'): 11,
	byte('A'): 12, byte('a'): 12,
	byte('T'): 13, byte('t'): 13,
	byte('E'): 14, byte('e'): 14,
}

func FrequencyAnalyzer(input []byte) []int {
	shiftedInput := make([]byte, 0)
	shiftedInput = append(shiftedInput, input[len(input)-1:]...)
	shiftedInput = append(shiftedInput, input[:len(input)-1]...)

	frequency := make([]int, 0)
	for i := 1; i < len(input)-1; i++ {
		coincidenceCount := 0
		for j := 0; j < len(input); j++ {
			if input[j] == shiftedInput[j] {
				coincidenceCount++
			}
		}

		frequency = append(frequency, coincidenceCount)
		shiftedInput = make([]byte, 0)
		shiftedInput = append(shiftedInput, input[len(input)-(i+1):]...)
		shiftedInput = append(shiftedInput, input[:len(input)-(i+1)]...)
	}
	return frequency
}

func RepeatingKeyXorCipher(input []byte, keyLength int) ([]byte, string) {
	subString := make([][]byte, keyLength)
	for i := 0; i < (len(input)/2)+1; i += keyLength {
		for j := 0; j < len(subString); j++ {
			if i+j >= len(input) {
				break
			}

			subString[j] = append(subString[j], input[j+i])
		}
	}

	key := ""
	for i := 0; i < len(subString); i++ {
		_, subKey := BruteForce(subString[i])
		key += subKey
	}

	return xor(input, []byte(key)), key
}

func xor(input, key []byte) []byte {
	var output []byte
	for i := 0; i < len(input); i++ {
		xorRes := input[i] ^ key[i%len(key)]
		output = append(output, xorRes)
	}

	return output
}

func totalDecryptScore(input []byte) int {
	totalScore := 0
	for i := 0; i < len(input); i++ {
		charScore, isOkay := charsScore[input[i]]
		if !isOkay {
			continue
		}

		totalScore += charScore
	}

	return totalScore
}

func possibleKeyValues() string {
	keyValues := ""
	for i := 48; i <= 122; i++ {
		if (i >= 58 && i <= 64) ||
			(i >= 91 && i <= 96) {
			continue
		}
		keyValues += string(byte(i))
	}

	return keyValues
}
