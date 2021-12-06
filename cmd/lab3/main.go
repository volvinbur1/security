package main

import (
	"fmt"
	"github.com/volvinbur1/security/internal/lab3/casinoroyale"
	"github.com/volvinbur1/security/internal/lab3/rng"
	"github.com/volvinbur1/security/internal/lab3/rng/lcg"
)

func main() {
	account := casinoroyale.NewAccount()
	lcgParams, err := crackLcg(account)
	if err != nil {
		panic(err)
	}

	account.SetLcgParameters(lcgParams)
	for account.Money < 1000000 {
		result, err := account.PlayLcg(account.Money)
		if err != nil {
			panic(err)
		}

		if !result.BetWon {
			fmt.Println("Opsss. Here a lose at lcg play.")
			break
		}
	}

	if account.Money >= 1000000 {
		fmt.Println("Yuhu, Victory at lcg play.")
	}
}

func crackLcg(account *casinoroyale.Account) (rng.Lcg, error) {
	var recentNumbers [3]int
	for i := 0; i < 3; i++ {
		result, err := account.PlayLcg(1)
		if err != nil {
			return rng.Lcg{}, err
		}

		recentNumbers[i] = int(result.RealNumber)
	}

	return lcg.CrackRng(recentNumbers)
}
