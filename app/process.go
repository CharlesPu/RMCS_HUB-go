package app

import (
	"RMCS_HUB/infra"
	"RMCS_HUB/protocol"
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
		infra.PrintHexList("process thread recv:", rxbuf)
		/* parse... */
		coId := rxbuf[6]
		staId := rxbuf[7]

		var bladeRTInfo infra.BladeRTInfo
		bladeRTInfo.CompanyID = int(coId)
		bladeRTInfo.StationID = int(staId)
		bladeRTInfo.Position = protocol.ByteToFloat32(rxbuf[8:12])

		var vibraPara infra.VibraPara
		vibraPara.CompanyID = int(coId)
		vibraPara.StationID = int(staId)
		vibraPara.BladeAm = protocol.ByteToFloat32(rxbuf[12:16])
		vibraPara.BladeEffCnt = protocol.ByteToUint32(rxbuf[16:20])
		vibraPara.CylinderAm = protocol.ByteToFloat32(rxbuf[20:24])
		vibraPara.CylinderEffCnt = protocol.ByteToUint32(rxbuf[24:28])
		vibraPara.AlarmNumber = protocol.ByteToUint16(rxbuf[28:30])

		cylinderRTInfo := make([]infra.CylinderRTInfo, 4)
		for i := 0; i < 4; i++ {
			cylinderRTInfo[i].CompanyID = int(coId)
			cylinderRTInfo[i].StationID = int(staId)
			cylinderRTInfo[i].CyID = i + 1
			cylinderRTInfo[i].Position = protocol.ByteToFloat32(rxbuf[(30 + i*4):(34 + i*4)])
		}
		/* store into MYSQL */
		ret := infra.InsertBladeRTInfoT(&bladeRTInfo)
		ret = infra.InsertVibrationParaT(&vibraPara)
		ret = infra.InsertCylinderRTInfoT(cylinderRTInfo)
		fmt.Println(ret)
		time.Sleep(100 * time.Millisecond)
	}
}

func Test() {
	fmt.Println("process test!")
}
