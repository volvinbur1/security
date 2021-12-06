package lcg

import (
	"fmt"
	"github.com/volvinbur1/security/internal/casinoroyale"
	"github.com/volvinbur1/security/internal/rng"
	"math"
)

func CrackRng(account *casinoroyale.Account) (rng.Lcg, error) {
	var recentNumbers []int
	for i := 0; i < 3; i++ {
		result, err := account.PlayLcg(1)
		if err != nil {
			return rng.Lcg{}, err
		}

		recentNumbers = append(recentNumbers, result.RealNumber)
	}

	m := int(math.Pow(2, 32))
	a := ((recentNumbers[2] - recentNumbers[1]) * ((recentNumbers[1] - recentNumbers[0]) % m)) % m
	c := (recentNumbers[1] - recentNumbers[0]*a) % m

	fmt.Println("Lcg rng cracked.")
	fmt.Println("a value:", a, "\tc value:", c)
	return rng.Lcg{
		A:          a,
		C:          a,
		LastNumber: recentNumbers[2],
	}, nil
}

func NextValue(lcgParams rng.Lcg) int {
	return (lcgParams.A*lcgParams.LastNumber + lcgParams.C) % int(math.Pow(2, 32))
}
