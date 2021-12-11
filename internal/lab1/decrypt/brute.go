package decrypt

func BruteForce(cipher []byte) ([]byte, string) {
	keyValues := possibleKeyValues()

	possibleDecryption := make(map[byte][]byte)
	for _, keyChar := range keyValues {
		decryption := xor(cipher, []byte{byte(keyChar)})
		possibleDecryption[byte(keyChar)] = decryption
	}

	maxScoredKey := byte(0)
	maxScoreValue := 0
	for key, bytes := range possibleDecryption {
		score := totalDecryptScore(bytes)
		if score > maxScoreValue {
			maxScoreValue = score
			maxScoredKey = key
		}
	}

	return possibleDecryption[maxScoredKey], string(maxScoredKey)
}
