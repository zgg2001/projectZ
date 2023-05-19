package transmission

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var MqttDataChan chan string

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Mqtt Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("Mqtt Connect lost: %v", err)
}

func MqttNewClient(spaceNum int) mqtt.Client {

	MqttDataChan = make(chan string, spaceNum)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", MqttServerIp, MqttServerPort))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername(MqttUsername)
	opts.SetPassword(MqttPassword)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panic("Error connect client", token.Error())
	}
	return client
}

func MqttDeleteClient(client mqtt.Client) {
	client.Disconnect(222)
	close(MqttDataChan)
}

// 发布
func MqttPub(client mqtt.Client, msg string) {
	topic := MqttWriteCmdTopic
	token := client.Publish(topic, 0, false, msg)
	token.Wait()
	log.Printf("Mqtt published msg to topic %s\n", topic)
}

// win发布
func WinMqttPub(client mqtt.Client, msg string) {
	topic := MqttWriteWinDataTopic
	token := client.Publish(topic, 0, false, msg)
	token.Wait()
	log.Printf("Mqtt published msg to topic %s\n", topic)
}

// 订阅
func MqttSub(client mqtt.Client) {
	topic := MqttReadDataTopic
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	log.Printf("Mqtt subscribed to topic %s\n", topic)
}

// 订阅消息处理
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	// log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	MqttDataChan <- string(msg.Payload())
}
