package infra

import (
	"fmt"
	"net"
	"sync"
)

const (
	DTUS_MAX_NUM = 256
)

type Dtu struct {
	connSock   net.Conn //need mutex!!!
	regPackInt int      // int(last 2 bytes and merge them)
	regPackHex [8]byte  //hex
	mu         sync.Mutex
}

var dtus []Dtu
var dtus_num rune = 0

func init() {
	fmt.Println("dtu.go init...")
	dtus = make([]Dtu, DTUS_MAX_NUM)
	for i := 0; i < DTUS_MAX_NUM; i++ {
		dtus[i].connSock = nil
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

func GetDTUId(co_id, sta_id uint32) int {
	return int(((co_id & 0x0f) << 4) | (sta_id & 0x0f))
}
func GetDTUFd(dtuId int) net.Conn {
	dtus[dtuId].mu.Lock()
	defer dtus[dtuId].mu.Unlock()
	return dtus[dtuId].connSock
}
func SetDTUFd(dtuId int, fd net.Conn) bool {
	dtus[dtuId].mu.Lock()
	defer dtus[dtuId].mu.Unlock()
	dtus[dtuId].connSock = fd
	return true
}

func Test() {
	fmt.Println("dtu test!")
}
