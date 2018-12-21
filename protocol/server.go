package protocol

import (
	"RMCS_HUB/infra"
	"bytes"
	"fmt"
	"net"
	"strconv"
	"time"
)

const (
	IP = "0.0.0.0"
	//IP         = "127.0.0.1"
	PORT       = 20001
	BUF_MAX    = 2048
	CLIENT_MAX = 256
)

var HUB_fd *net.TCPListener = nil

func init() {
	fmt.Println("server.go init...")
	service := IP + ":" + strconv.Itoa(PORT)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	infra.CheckError(err)
	HUB_fd, err = net.ListenTCP("tcp4", tcpAddr)
}

func Receive(que chan<- []byte) {
	for {
		connFd, err := HUB_fd.Accept()
		if err != nil {
			infra.CheckError(err)
			continue
		}
		fmt.Printf("client %d has connected\n", connFd)
		go receiveHandle(connFd, que) //go!
		time.Sleep(1000 * time.Millisecond)
	}
}

func receiveHandle(connFd net.Conn, que chan<- []byte) int {
	defer connFd.Close()
	/* first receive */
	rxbuf := make([]byte, BUF_MAX)
	dtuId := -1
	for {
		bytesNum, _ := connFd.Read(rxbuf)
		infra.PrintHexList("srv recv:", rxbuf[:bytesNum])
		if bytesNum >= 8 { // or == 50
			/* get reg pack*/
			fmt.Printf("reg_pack: %02x\n", rxbuf[:8])
			/* check reg pack in mysql*/
			if infra.ExistDtu(rxbuf[:8]) { // exist
				//update this dtu's fd
				dtuId = int(((rxbuf[6] & 0x0f) << 4) | (rxbuf[7] & 0x0f))
				infra.SetDTUFd(dtuId, connFd)
				/* parse... */
				frameHeader := rxbuf[8:12]
				xorVal := GenerateXorValue(rxbuf[8:50])
				if bytes.Equal(frameHeader, FRAME_HEADER[:]) && xorVal == rxbuf[50] { // correct frame
					rxbuf = append(rxbuf[:8], rxbuf[12:50]...) // len = 50 - 4
					/* put into channel */
					que <- rxbuf
				}
			} else {
				fmt.Println("illegal DTU!")
			}
		} else if bytesNum <= 0 {
			fmt.Printf("client %d has closed the connection\n", connFd)
			//find dtu and reset fd
			if dtuId != -1 {
				infra.SetDTUFd(dtuId, nil)
			}
			return -1
		}
		time.Sleep(10 * time.Millisecond)
	}

	return 0
}

func Test() {
	fmt.Println("server test!")
}
