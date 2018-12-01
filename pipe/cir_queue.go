package pipe

import (
	"RMCS_HUB/infra"
	"fmt"
)

const (
	CQ_BUF_MAX  = 512
	CQ_CELL_MAX = 512
)

type qCell struct {
	d     *(infra.Dtu)
	q_buf []byte
}

var (
	head int = 0
	tail int = 0
	cq   []qCell
)

func init() {
	cq = make([]qCell, CQ_BUF_MAX)
	for i := 0; i < CQ_BUF_MAX; i++ {
		cq[i].d = nil
		cq[i].q_buf = make([]byte, CQ_CELL_MAX)
	}
}

func IsCQEmpty() int {
	if head == tail {
		return 0
	} else {
		return 1
	}
}

func IsCQFull() int {
	leftNum := (tail - head - 1) & (CQ_BUF_MAX - 1)
	if leftNum == 0 {
		return 0
	} else {
		return 1
	}
}

func GetCell(n int) ([]byte, *(infra.Dtu)) {
	tail = (tail + 1) & (CQ_BUF_MAX - 1)

	return cq[tail].q_buf[:n], cq[tail].d
}

func PutCell(_d *(infra.Dtu), b []byte, n int) (rn int) {
	cq[head].d = _d
	for i := 0; i < CQ_CELL_MAX; i++ {
		cq[head].q_buf[i] = 0
	}
	rn = copy(cq[head].q_buf, b)

	head = (head + 1) & (CQ_BUF_MAX - 1)

	return rn
}

func Test() {
	fmt.Println("cir_queue test!")
}
