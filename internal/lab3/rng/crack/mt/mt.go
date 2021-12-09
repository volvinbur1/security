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
		}
	}

	if seed == 0 {
		return errors.New("seed value not found for mt19937 algorithm")
	}

	account.SeedMtRandom(seed)
	fmt.Println("Mt19937 cracked. Seed timestamp", time.Unix(int64(seed), 0).Format(time.RFC3339))
	return nil
}

func RecoverStates(account *casinoroyale.Account) error {
	var states []uint32
	fmt.Println("Recovering states...")
	for i := 0; i < mt19937.N; i++ {
		result, err := account.PlayMtBetter(1)
		if err != nil {
			return err
		}

		states = append(states, invertTemper(uint32(result.RealNumber)))
	}

	account.SeedMtBetterRandom(states)
	fmt.Println("States recovered. Mt19937 with better seed cracked.")
	return nil
}

func invertTemper(val uint32) uint32 {
	val ^= val >> mt19937.L
	val ^= (val << mt19937.T) & mt19937.C

	for i := 0; i < 4; i++ {
		b := uint32(mt19937.B) & uint32(0x7f<<(mt19937.S*(i+1)))
		val ^= (val << mt19937.S) & b
	}

	for i := 0; i < 3; i++ {
		val ^= val >> mt19937.U
	}

	return val
}
