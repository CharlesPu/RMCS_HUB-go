package main

import (
	"RMCS_HUB/app"
	// "RMCS_HUB/infra"
	// "RMCS_HUB/pipe"
	"RMCS_HUB/protocol"
	"fmt"
)

const (
	VERSION = "1.0.0"
	LANG    = "golang"
	QSize   = 256
)

func init() {
	fmt.Println("main init...")
}
func main() {
	// infra.ShowAllTables()

	queueChan := make(chan []byte, QSize)
	go protocol.Receive(queueChan)
	go app.Process(queueChan)
	go app.CMDSend()

	for {
	}
}
