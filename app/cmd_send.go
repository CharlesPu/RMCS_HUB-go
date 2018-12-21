package app

import (
	"RMCS_HUB/infra"
	"RMCS_HUB/protocol"
	"fmt"
	"time"
)

func init() {
	fmt.Println("cmd_send.go init...")
}

func CMDSend() {
	// truncate ctrl tables
	infra.TruncateTables("VibrationControlPara", "SingleCyControlPara")
	for {
		var vc infra.VibraCtrl
		vc, ret := infra.ReadVibrationControlPara()
		if ret == true { //package
			txbuf := make([]byte, 31)
			// txbuf = append(txbuf[:0], infra.FRAME_HEADER[:]...)
			copy(txbuf, protocol.FRAME_HEADER[:])
			txbuf[4] = protocol.TYPE_SYSPARA
			// if use append(), len(txbuf) will not be 31 after each append(),
			// and will be the length of what append() returns
			copy(txbuf[5:], protocol.Float32ToByte(vc.BladeAm))
			copy(txbuf[9:], protocol.Uint32ToByte(vc.BladeVibraCnt))
			copy(txbuf[13:], protocol.Float32ToByte(vc.CylinderAm))
			copy(txbuf[17:], protocol.Uint32ToByte(vc.CylinderCycle))
			// send
			coId := vc.CompanyID
			staId := vc.StationID
			dtuId := ((coId & 0x0f) << 4) | (staId & 0x0f)
			if fd := infra.GetDTUFd(dtuId); fd != nil { // online
				txbuf[30] = protocol.GenerateXorValue(txbuf[:30])
				infra.PrintHexList("cmd_send:", txbuf)
				fd := infra.GetDTUFd(dtuId)
				// if fd != nil {
				fd.Write(txbuf[:31])
				// }
			} //else drop
		}

		var scc []infra.SingleCyCtrl
		scc, num, ret := infra.ReadSingleCyControlPara()
		if ret == true && num != 0 && scc != nil {
			for _, val := range scc { //package each one
				txbuf := make([]byte, 31)
				// txbuf = append(txbuf[:0], infra.FRAME_HEADER[:]...)
				copy(txbuf, protocol.FRAME_HEADER[:])
				if val.JogUp != 0 { //or != 0?
					txbuf[4] = protocol.TYPE_CY_JOGUP
					txbuf[21] = 1 << uint32(val.CyID-1)
				} else if val.JogDown != 0 {
					txbuf[4] = protocol.TYPE_CY_JOGDOWN
					txbuf[21] = 1 << uint32(val.CyID-1)
				} else if val.Reset != 0 {
					txbuf[4] = protocol.TYPE_CY_RESET
					txbuf[21] = 1 << uint32(val.CyID-1)
				} else if val.Amplitude != 0 {
					txbuf[4] = protocol.TYPE_CY_PARA
					txbuf[21] = 1 << uint32(val.CyID-1)
					// if use append(), len(txbuf) will not be 31 after each append(),
					// and will be the length of what append() returns
					copy(txbuf[22:], protocol.Float32ToByte(val.Amplitude))
					copy(txbuf[26:], protocol.Uint32ToByte(val.Cycle))
				}
				//send
				coId := val.CompanyID
				staId := val.StationID
				dtuId := ((coId & 0x0f) << 4) | (staId & 0x0f)
				if fd := infra.GetDTUFd(dtuId); fd != nil { // online
					txbuf[30] = protocol.GenerateXorValue(txbuf[:30])
					infra.PrintHexList("cmd_send:", txbuf)
					fd := infra.GetDTUFd(dtuId)
					// if fd != nil {
					fd.Write(txbuf[:31])
					// }
				} //else drop
			}
		}

		time.Sleep(1000 * time.Millisecond)
	}
}
