package infra

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("errors.go init...")
}

func CheckError(err error) int {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		return 1
	}
	return 0
}
