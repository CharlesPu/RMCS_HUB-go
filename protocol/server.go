package protocol

import (
	"RMCS_HUB/infra"
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
	// fmt.Println(service)
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
	for {
		bytesNum, _ := connFd.Read(rxbuf)
		infra.PrintHexList("srv recv:", rxbuf[:bytesNum])
		if bytesNum >= 8 {
			/* get reg pack*/
			fmt.Printf("reg_pack: %02x\n", rxbuf[:8])
			/* check reg pack in mysql*/
			if infra.ExistDtu(rxbuf[:8]) { // exist
				/* parse... */

				/* put into channel */
				que <- rxbuf[:bytesNum]
			} else {
				fmt.Println("illegal DTU!")
			}
		} else if bytesNum <= 0 {
			fmt.Printf("client %d has closed the connection\n", connFd)
			return -1
		}
		time.Sleep(1 * time.Millisecond)
	}

	return 0
}

func Test() {
	fmt.Println("server test!")
}
