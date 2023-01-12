package main

import (
	"raspberry/internal/pi"
	"raspberry/internal/transmission"
	"time"
)

func main() {

	// 管道初始化
	pi.PipeInit()
	defer pi.PipeRemove()

	// 硬件脚本执行/交互
	cmd := pi.PythonStartUp()
	defer pi.PythonCancel(cmd)
	go pi.PythonRunTask()

	// mqtt client
	mqttClient := transmission.MqttNewClient()
	defer transmission.MqttDeleteClient(mqttClient)

	// mqtt 订阅启动
	transmission.MqttSub(mqttClient)

	time.Sleep(time.Second * 1000)
}
