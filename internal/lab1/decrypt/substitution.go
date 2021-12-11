package decrypt

func Substitution(cipher, key []byte) []byte {
	var plaintext []byte
	for i := 0; i < len(cipher); i++ {
		plaintext = append(plaintext, key[cipher[i]-'A'])
	}

	return plaintext
}
