package infra

import (
	"fmt"
)

func init() {
	fmt.Println("logs.go init...")
}

func PrintHexList(s string, l []byte) {
	fmt.Printf("%s(%d)\n", s, len(l))
	for i := range l {
		fmt.Printf("%02x ", l[i])
	}
	fmt.Printf("\n")
}
