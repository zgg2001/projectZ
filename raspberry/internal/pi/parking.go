package pi

import (
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/zgg2001/projectZ/raspberry/internal/transmission"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type status = int

type ParkingOperation interface {
	Reset(id int)
	GetStatus() status
	DriveInto(license string, cli *mqtt.Client)
	DriveOut(cli *mqtt.Client)
	UpdataData(dataStr string) error
}

type Parking struct {
	Id      int
	IsUsing bool
	License string

	Temperature int32
	Humidity    int32
	IsFlame     int32
	IsFlammable int32
}

func (p *Parking) GetStatus() status {
	if p.IsUsing {
		return NonEmptyParkingSpace
	}
	return EmptyParkingSpace
}

func (p *Parking) Reset(id int) {
	p.Id = id
	p.IsUsing = false
	p.License = ""
}

func (p *Parking) DriveInto(license string, cli *mqtt.Client) {
	transmission.MqttPub(*cli, fmt.Sprintf("%d:%d", p.Id, transmission.CmdServoUp))
	p.IsUsing = true
	p.License = license
}

func (p *Parking) DriveOut(cli *mqtt.Client) {
	transmission.MqttPub(*cli, fmt.Sprintf("%d:%d", p.Id, transmission.CmdServoDown))
	p.IsUsing = false
	p.License = ""
}

func (p *Parking) UpdataData(strArr []string) error {

	temperature, err := strconv.ParseInt(strArr[1], 10, 64)
	if err != nil {
		return err
	}
	humidity, err := strconv.ParseInt(strArr[2], 10, 64)
	if err != nil {
		return err
	}
	isFlame, err := strconv.ParseInt(strArr[3], 10, 64)
	if err != nil {
		return err
	}
	isFlammable, err := strconv.ParseInt(strArr[4], 10, 64)
	if err != nil {
		return err
	}

	atomic.StoreInt32(&p.Temperature, int32(temperature))
	atomic.StoreInt32(&p.Humidity, int32(humidity))
	atomic.StoreInt32(&p.IsFlame, int32(isFlame))
	atomic.StoreInt32(&p.IsFlammable, int32(isFlammable))
	return nil
}

func (p *Parking) GetData() (int32, int32, int32) {

	temperature := atomic.LoadInt32(&p.Temperature)
	humidity := atomic.LoadInt32(&p.Humidity)
	isFlame := atomic.LoadInt32(&p.IsFlame)
	isFlammable := atomic.LoadInt32(&p.IsFlammable)

	alarm := NoAlarm
	if isFlame == 1 {
		alarm += FireAlarm
	}
	if isFlammable == 1 {
		alarm += GasAlarm
	}

	return temperature, humidity, alarm
}
