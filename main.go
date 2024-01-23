package main

import (
	"fmt"

	"zp.com/wallet-sdk/common"
	"zp.com/wallet-sdk/sdk"
)

func main() {
	fmt.Printf("\x1b[%dm%s\x1b[0m %s\n", common.Blue, "tag:       ", common.Tag)
	fmt.Printf("\x1b[%dm%s\x1b[0m %s\n", common.Blue, "commit:    ", common.Commit)
	fmt.Printf("\x1b[%dm%s\x1b[0m %s\n", common.Blue, "build time:", common.BuildTime)

	btcAddress := sdk.BTCNewAddress("0357bbb2d4a9cb8a2357633f201b9c518c2795ded682b7913c6beef3fe23bd6d2f")
	fmt.Println("zp:::" + btcAddress)
}
