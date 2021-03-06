package main

import (
	"fmt"
	"github.com/volvinbur1/security/internal/lab3/casinoroyale"
	"github.com/volvinbur1/security/internal/lab3/rng/crack/lcg"
	"github.com/volvinbur1/security/internal/lab3/rng/crack/mt"
	"os"
)

type WorkMode int

const (
	lcgPlay WorkMode = iota
	mtPlay
	betterMtPlay
)

func main() {
	fmt.Println("-----------------Linear congruential generator-----------------")
	playCasino(lcgPlay)
	fmt.Println("---------------------Mersenne Twister 19937--------------------")
	playCasino(mtPlay)
	fmt.Println("--------------Mersenne Twister 19937 (Better seed)-------------")
	playCasino(betterMtPlay)
}

func getWorkingMode() WorkMode {
	if len(os.Args) < 2 {
		panic("work mode not specified")
	}

	switch os.Args[1] {
	case "--lcg":
		fmt.Println("-----------------Linear congruential generator-----------------")
		return lcgPlay
	case "--mt":
		fmt.Println("---------------------Mersenne Twister 19937--------------------")
		return mtPlay
	case "--mtb":
		fmt.Println("--------------Mersenne Twister 19937 (Better seed)-------------")
		return betterMtPlay
	default:
		panic("unknown working mode")
	}
}

func playCasino(workMode WorkMode) {
	account, err := createCasinoAccount(workMode)
	if err != nil {
		panic(err)
	}

	for account.Money < 1000000 {
		var result casinoroyale.PlayResult
		switch workMode {
		case lcgPlay:
			result, err = account.PlayLcg(account.Money)
			break
		case mtPlay:
			result, err = account.PlayMt(account.Money)
			break
		case betterMtPlay:
			result, err = account.PlayMtBetter(account.Money)
			break
		}

		if err != nil {
			panic(err)
		}

		if !result.BetWon {
			fmt.Println("Opsss. Here a lose.")
			break
		}
	}

	if account.Money >= 1000000 {
		fmt.Println("Yuhu, Victory.")
	}
}

func createCasinoAccount(mode WorkMode) (*casinoroyale.Account, error) {
	account := casinoroyale.NewAccount()

	switch mode {
	case lcgPlay:
		err := crackLcg(account)
		if err != nil {
			return nil, err
		}
		break
	case mtPlay:
		err := mt.RecoverSeed(account)
		if err != nil {
			return nil, err
		}
		break
	case betterMtPlay:
		err := mt.RecoverStates(account)
		if err != nil {
			return nil, err
		}
		break
	}

	return account, nil
}

func crackLcg(account *casinoroyale.Account) error {
	var recentNumbers [3]int
	for i := 0; i < 3; i++ {
		result, err := account.PlayLcg(1)
		if err != nil {
			return err
		}

		recentNumbers[i] = int(result.RealNumber)
	}

	lcgParams, err := lcg.CrackRng(recentNumbers)
	if err != nil {
		return err
	}

	account.SetLcgParameters(lcgParams)
	return nil
}
