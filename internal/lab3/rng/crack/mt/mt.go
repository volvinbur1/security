package mt

import (
	"errors"
	"fmt"
	"github.com/volvinbur1/security/internal/lab3/casinoroyale"
	"github.com/volvinbur1/security/internal/lab3/rng/mt19937"
	"time"
)

const bounds = 15

func RecoverSeed(account *casinoroyale.Account) error {
	startTime := time.Now()
	result, err := account.PlayMt(1)
	if err != nil {
		return err
	}

	seed := uint32(0)
	for timestamp := startTime.UTC().Unix() - bounds; timestamp <= startTime.UTC().Unix()+bounds; timestamp++ {
		rand := mt19937.New(uint32(timestamp))
		rv := rand.UInt32()
		if rv == uint32(result.RealNumber) {
			seed = uint32(timestamp)
			fmt.Println("Mt19937 cracked. Seed timestamp", time.Unix(timestamp, 0).Format(time.RFC3339))
		}
	}

	if seed == 0 {
		return errors.New("seed value not found for mt19937 algorithm")
	}

	account.SeedMtRandom(seed)
	return nil
}
