package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	work_dir           = "./tmp"
	pipe_read          = "./tmp/pipe.1"
	pipe_write         = "./tmp/pipe.2"
	MQTT_SERVER_IP     = "192.168.53.133"
	MQTT_SERVER_PORT   = 1883
	MQTT_READ_USERNAME = "test2"
	MQTT_READ_PASSWORD = "b123456"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Printf("Connect lost: %v", err)
}

func main() {
	// tmp目录创建
	if _, err := os.Stat(work_dir); err != nil {
		err := os.MkdirAll(work_dir, 0711)
		if err != nil {
			log.Println("Error creating directory: ", err)
			return
		}
	}

	// 管道初始化
	os.Remove(pipe_read)
	os.Remove(pipe_write)
	err := syscall.Mkfifo(pipe_read, 0666)
	if err != nil {
		log.Println("Error make read fifo", err)
	}
	defer os.Remove(pipe_read)
	err = syscall.Mkfifo(pipe_write, 0666)
	if err != nil {
		log.Println("Error make write fifo", err)
	}
	defer os.Remove(pipe_write)

	/*
		// mqtt服务架设
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
		defer client.Disconnect(250)

		// 订阅
		sub(client)
	*/

	// 硬件数据交互
	go python_task()

	// 硬件逻辑
	cmd := exec.Command("/usr/bin/python3", "./script/raspberry.py")
	err = cmd.Start()
	if err != nil {
		log.Println("Error start python script", err)
	}
	defer cmd.Process.Signal(syscall.SIGQUIT)

	log.Println("Start work")
	for {

	}
}

// 订阅
func sub(client mqtt.Client) {
	topic := "my/mqtt/topic"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	log.Printf("Subscribed to topic %s\n", topic)
}

// 消息处理
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

// 树莓派硬件数据交互
func python_task() {
	r_pipe, err := os.OpenFile(pipe_read, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		log.Println("Error open file: ", err)
	}
	w_pipe, err := os.OpenFile(pipe_write, os.O_RDWR, 0777)
	if err != nil {
		log.Println("Error open file: ", err)
	}

	reader := bufio.NewReader(r_pipe)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println("Error read bytes: ", err)
		}
		plate_str := strings.TrimRight(string(line), "\n")
		log.Println(plate_str)

		// 检测逻辑部分
		if plate_str == "沪AZ0001" {
			w_pipe.WriteString("pass")
		} else {
			w_pipe.WriteString("reject")
		}
	}
}
