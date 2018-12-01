package infra

import (
	"fmt"
)

func init(){
	fmt.Println("logs.go init...")
}

func PrintHexList(s string, l []byte) {
	fmt.Println(s)
	for i, _ := range l {
		fmt.Printf("%02x ", l[i])
	}
	fmt.Printf("\n")
}
