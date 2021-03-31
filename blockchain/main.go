package main

import (
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

type Network struct {
	name        string
	symbol      string
	xpubkey     byte
	xprivatekey byte
}

var network = map[string]Network{
	"btc": {name: "bitcoin", symbol: "btc", xpubkey: 0x00, xprivatekey: 0x80},
	"ltc": {name: "litecoin", symbol: "ltc", xpubkey: 0x30, xprivatekey: 0xb0},
}

func (network Network) GetnetworkParams() *chaincfg.Params {
	networkParams := &chaincfg.MainNetParams
	networkParams.PubKeyHashAddrID = network.xpubkey
	networkParams.PrivateKeyID = network.xprivatekey
	return networkParams
}

func (network Network) CreatePrivateKey() (*btcutil.WIF, error) {
	secret, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}

	return btcutil.NewWIF(secret, network.GetnetworkParams(), true)
}

func (network Network) GetAddress(wif *btcutil.WIF) (*btcutil.AddressPubKey, error) {
	return btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), network.GetnetworkParams())
}

func (network Network) ImportWIF(wifStr string) (*btcutil.WIF, error) {
	wif, err := btcutil.DecodeWIF(wifStr)
	if err != nil {
		return nil, err
	}

	if !wif.IsForNet(network.GetnetworkParams()) {
		return nil, errors.New("The WIF string is not valid for the `" + network.name +"` network")
	}
	return wif, nil
}

func main() {
	wif, _ := network["btc"].CreatePrivateKey()
	address, _ := network["btc"].GetAddress(wif)
	fmt.Printf("Bitcoin WIF: %s - Address: %s \n", wif.String(), address.EncodeAddress())
	wif, _ = network["ltc"].CreatePrivateKey()
	address, _ = network["ltc"].GetAddress(wif)
	fmt.Printf("Litecoin WIF: %s - Address: %s", wif.String(), address.EncodeAddress())
}