package decrypt

func BruteForce(cipher []byte, possibleKeyValues string) []byte {
	possibleDecryption := make(map[byte][]byte)
	for _, keyChar := range possibleKeyValues {
		decryption := xor(cipher, byte(keyChar))
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

	return possibleDecryption[maxScoredKey]
}
