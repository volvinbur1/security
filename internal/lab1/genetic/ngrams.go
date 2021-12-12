package genetic

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func ngrams(str []byte, n int) map[string]float64 {
	ngramsMap := make(map[string]float64)

	for i := 0; i < len(str)-n+1; i++ {
		substr := str[i : i+n]
		_, isOkay := ngramsMap[string(substr)]
		if isOkay {
			ngramsMap[string(substr)]++
		} else {
			ngramsMap[string(substr)] = 1
		}
	}

	for key, val := range ngramsMap {
		ngramsMap[key] = val / float64(len(str)-n+1)
	}

	return ngramsMap
}

func parseTrigramsFile(filepath string) (map[string]float64, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	trigrams := make(map[string]float64)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) != 2 {
			return nil, errors.New("error in trigrams file format")
		}

		trigrams[parts[0]], err = strconv.ParseFloat(parts[1], 64)
		if err != nil {
			return nil, err
		}
	}

	return trigrams, err
}
