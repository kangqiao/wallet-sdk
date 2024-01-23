package sdk

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/okx/go-wallet-sdk/coins/bitcoin"
)

func BTCNewAddress(pubKey string) string {
	// address
	network := &chaincfg.TestNet3Params
	pubKeyHex := pubKey
	publicKey, _ := hex.DecodeString(pubKeyHex)
	p2pkh, err := bitcoin.PubKeyToAddr(publicKey, bitcoin.LEGACY, network)
	if err != nil {
		// todo
		fmt.Println(err)
		return ""
	}
	fmt.Println(p2pkh)
	return p2pkh
}
