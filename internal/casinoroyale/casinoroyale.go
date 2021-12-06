package casinoroyale

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/volvinbur1/security/internal/rng"
	"github.com/volvinbur1/security/internal/rng/lcg"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const baseUrl = "http://95.217.177.249/casino"

// endpoints
const (
	createAcc = "createacc"
	playLcg   = "playLcg"
	//playMt       = "playMt"
	//playBetterMt = "playBetterMt"
)

type Account struct {
	Id    string `json:"id"`
	Money int    `json:"money"`

	isLcgCracked  bool
	lcgParameters rng.Lcg
}

type PlayResult struct {
	BetWon     bool
	Message    string
	RealNumber int
}

func NewAccount() *Account {
	account := Account{}

	rand.Seed(time.Now().UnixNano())
	randomIdValue := rand.Intn(2000) + 1000
	for {
		body, err := requestToCasino(fmt.Sprintf("%s/%s?id=%d", baseUrl, createAcc, randomIdValue))
		if err != nil {
			if body == nil {
				panic(err)
			}

			randomIdValue = rand.Intn(2000) + 1000
			continue
		}

		err = json.Unmarshal(body, &account)
		if err != nil {
			panic(err)
		}

		fmt.Println("Account created at Casino Royale.")
		fmt.Println("Account id", account.Id, "Starting balance", account.Money)
		break
	}

	return &account
}

func (a *Account) SetLcgParameters(lcgParams rng.Lcg) {
	a.isLcgCracked = true
	a.lcgParameters = lcgParams
}

func (a *Account) PlayLcg(betAmount int) (PlayResult, error) {
	betNumber := 1
	if a.isLcgCracked {
		betNumber = lcg.NextValue(a.lcgParameters)
	}

	body, err := requestToCasino(fmt.Sprintf("%s/%s?id=%s&bet=%d&number=%d",
		baseUrl, playLcg, a.Id, betAmount, betNumber))
	if err != nil {
		return PlayResult{}, err
	}

	result, err := parsePlayResponse(body)
	if err != nil {
		return PlayResult{}, err
	}

	a.lcgParameters.LastNumber = result.RealNumber
	if result.RealNumber == betNumber {
		a.Money += betAmount
	} else {
		a.Money -= betAmount
	}

	return result, nil
}

func parsePlayResponse(body []byte) (PlayResult, error) {
	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		return PlayResult{}, err
	}

	msg, isOkay := data["message"].(string)
	if !isOkay {
		return PlayResult{}, errors.New("key `message` not found")
	}

	realNumber, isOkay := data["realNumber"].(float64)
	if !isOkay {
		return PlayResult{}, errors.New("key `realNumber` not found")
	}

	return PlayResult{
		Message:    msg,
		RealNumber: int(realNumber),
	}, nil
}

func requestToCasino(requestUrl string) ([]byte, error) {
	resp, err := http.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode/100 != 2 {
		casinoErr := struct{ error string }{}
		err = json.Unmarshal(body, &casinoErr)
		if err != nil {
			return nil, err
		}

		return body, errors.New(fmt.Sprint(resp.Status, ". ", casinoErr.error))
	}

	return body, nil
}
