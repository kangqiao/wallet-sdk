package main

import (
	"fmt"

	"zp.com/wallet-sdk/common"
)

func main() {
	fmt.Printf("\x1b[%dm%s\x1b[0m %s\n", common.Blue, "tag:       ", common.Tag)
	fmt.Printf("\x1b[%dm%s\x1b[0m %s\n", common.Blue, "commit:    ", common.Commit)
	fmt.Printf("\x1b[%dm%s\x1b[0m %s\n", common.Blue, "build time:", common.BuildTime)

	fmt.Println("zhaopan")
}
