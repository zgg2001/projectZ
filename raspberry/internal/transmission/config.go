package transmission

const (
	MqttServerIp      = "192.168.242.133"
	MqttServerPort    = 1883
	MqttUsername      = "test0"
	MqttPassword      = "z123456"
	MqttReadDataTopic = "pi/esp32/data"
	MqttWriteCmdTopic = "pi/esp32/cmd"

	RPCServerIddr = "xxx.xxx.xxx.xxx:8888"
	ClientPemPath = "./auth/client.pem"
	ClientKeyPath = "./auth/client.key"
	CACrtPath     = "./auth/ca.crt"
)

//cmd number
var (
	CmdInvalid   int = 0
	CmdServoUp   int = 1
	CmdServoDown int = 2
)
