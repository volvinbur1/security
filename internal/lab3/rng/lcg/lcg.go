package lcg

import (
	"fmt"
	"github.com/volvinbur1/security/internal/lab3/rng"
	"math"
	"math/big"
)

func CrackRng(recentNumbers [3]int) (rng.Lcg, error) {
	m := int64(math.Pow(2, 32))

	mul := big.NewInt(int64(recentNumbers[1] - recentNumbers[0]))
	mul.ModInverse(mul, big.NewInt(m))
	mul.Mul(mul, big.NewInt(int64(recentNumbers[2]-recentNumbers[1])))
	a := int(mul.Mod(mul, big.NewInt(m)).Int64())

	c := int32((recentNumbers[1] - recentNumbers[0]*a) % int(m))

	fmt.Print("Lcg rng cracked. ")
	fmt.Println("a value:", a, "\tc value:", c)
	return rng.Lcg{
		A:          int32(a),
		C:          c,
		LastNumber: int32(recentNumbers[2]),
	}, nil
}

func NextValue(lcgParams rng.Lcg) int32 {
	return (lcgParams.A*lcgParams.LastNumber + lcgParams.C) % int32(math.Pow(2, 32))
}
