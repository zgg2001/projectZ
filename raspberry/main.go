package main

import (
	"github.com/zgg2001/projectZ/raspberry/internal/pi"
	"github.com/zgg2001/projectZ/raspberry/internal/transmission"

	"time"
)

var parkingSpacesNum int = 2

func main() {

	pi.PipeInit()
	defer pi.PipeRemove()

	mqttClient := transmission.MqttNewClient(parkingSpacesNum)
	defer transmission.MqttDeleteClient(mqttClient)

	var mgr pi.ParkingMgr
	mgr.Init(parkingSpacesNum, &mqttClient)

	cmd := pi.PythonStartUp()
	defer pi.PythonCancel(cmd)

	go pi.RunPythonTask(&mgr)
	go pi.RunDataTask(transmission.MqttDataChan, &mgr)

	// mqtt 订阅启动
	transmission.MqttSub(mqttClient)

	//
	pi.UploadPiData(&mgr)

	time.Sleep(time.Second * 1000)
}
