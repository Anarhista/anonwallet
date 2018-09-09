package main

import (
	"encoding/json"
	"log"
	"os/exec"
	"strconv"
)

type BitcoinGenerator struct {
}

func (bg *BitcoinGenerator) getAddress() (string, error) {
	out, err := exec.Command("electrum", "createnewaddress").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (bg *BitcoinGenerator) getBalance(address string) (float64, error) {
	out, err := exec.Command("/usr/local/bin/electrum", "getaddressbalance", address).CombinedOutput()
	if err != nil {
		log.Println(string(out))
		return 0, err
	}

	br := &BalanceResponse{}
	err = json.Unmarshal([]byte(out), br)
	if err != nil {
		return 0, err
	}

	balance, err := strconv.ParseFloat(br.Confirmed, 64)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func initBtcGen() *BitcoinGenerator {
	bg := &BitcoinGenerator{}
	return bg
}

type BalanceResponse struct {
	Confirmed   string `json:"confirmed"`
	Unconfirmed string `json:"unconfirmed"`
}
