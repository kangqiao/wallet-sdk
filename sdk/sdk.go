package sdk

import "fmt"

func PreSign(r string) string {
	fmt.Println("pre sign")
	return "succeed"
}

func Sign(r string) string {
	fmt.Println("pre sign")
	return "result"
}

func Verify(r string) bool {
	fmt.Println("verify")
	return true
}
