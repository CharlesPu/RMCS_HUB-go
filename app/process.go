package app

import (
	"RMCS_HUB/infra"
	"fmt"
	"time"
)

func init() {
	fmt.Println("process.go init...")
}

func Process(que <-chan []byte) {
	for {
		/* get from channel */
		rxbuf := <-que
		length := len(rxbuf)
		infra.PrintHexList("process thread recv:", rxbuf[:length])
		/* parse... */

		/* store into MYSQL */

		time.Sleep(10 * time.Millisecond)
	}
}

func Test() {
	fmt.Println("process test!")
}
