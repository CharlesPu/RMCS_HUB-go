package infra

import (
	"fmt"
	"net"
)

const (
	DTUS_MAX_NUM = 256
)

type Dtu struct {
	dtuSock    int //need mutex!!!
	connSock   net.Conn
	regPackInt int     // int(last 2 bytes and merge them)
	regPackHex [8]byte //hex
}

var dtus []Dtu
var dtus_num rune = 0

func init() {
	fmt.Println("dtu.go init...")
	dtus = make([]Dtu, DTUS_MAX_NUM)
	for i := 0; i < DTUS_MAX_NUM; i++ {
		dtus[i].regPackInt = i
		dtus[i].regPackHex[0] = 0x77
		dtus[i].regPackHex[1] = 0x77
		dtus[i].regPackHex[2] = 0x77
		dtus[i].regPackHex[3] = 0x77
		dtus[i].regPackHex[4] = 0x00
		dtus[i].regPackHex[5] = 0x00
		dtus[i].regPackHex[6] = byte((i >> 4) & 0x0f)
		dtus[i].regPackHex[7] = byte(i & 0x0f)
		// fmt.Println(dtus[i].regPackHex, i)
	}
}

func Test() {
	fmt.Println("dtu test!")
}
