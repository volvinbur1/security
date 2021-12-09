package mt

import (
	"fmt"
	"github.com/volvinbur1/security/internal/lab3/rng/mt19937"
	"time"
)

const bounds = 15

func CrackRng(realNumber uint32, startTime time.Time) uint32 {
	for val := startTime.UTC().Unix() - bounds; val <= startTime.UTC().Unix()+bounds; val++ {
		rand := mt19937.New(uint32(val))
		rv := rand.UInt32()
		if rv == realNumber {
			fmt.Println("Mt19937 cracked. Seed timestamp", time.Unix(val, 0).Format(time.RFC3339))
			return uint32(val)
		}
	}

	return 0
}
