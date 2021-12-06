package casinoroyale

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const baseUrl = "http://95.217.177.249/casino"

// endpoints
const (
	createAcc    = "createacc"
	playLcg      = "playLcg"
	playMt       = "playMt"
	playBetterMt = "playBetterMt"
)

type Account struct {
	Id    string `json:"id"`
	Money int    `json:"money"`
}

func NewAccount() Account {
	account := Account{}

	rand.Seed(time.Now().UnixNano())
	randomIdValue := rand.Intn(2000) + 1000
	for {
		resp, err := http.Get(fmt.Sprintf("%s/%s?id=%d", baseUrl, createAcc, randomIdValue))
		if err != nil {
			panic(err)
		}

		if resp.StatusCode/100 != 2 {
			randomIdValue = rand.Intn(2000) + 1000
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(body, &account)
		if err != nil {
			panic(err)
		}
		break
	}

	return account
}
