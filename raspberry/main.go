package main

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	MQTT_SERVER_IP     = "192.168.53.133"
	MQTT_SERVER_PORT   = 1883
	MQTT_READ_USERNAME = "test2"
	MQTT_READ_PASSWORD = "b123456"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", MQTT_SERVER_IP, MQTT_SERVER_PORT))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername(MQTT_READ_USERNAME)
	opts.SetPassword(MQTT_READ_PASSWORD)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	sub(client)

	for {

	}

	client.Disconnect(250)
}

//订阅
func sub(client mqtt.Client) {
	topic := "my/mqtt/topic"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s\n", topic)
}

//消息处理
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
