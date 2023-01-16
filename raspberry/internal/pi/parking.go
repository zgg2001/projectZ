package pi

import (
	"fmt"
	"raspberry/internal/transmission"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type status = int

type ParkingOperation interface {
	Reset(id int)
	GetStatus() status
	DriveInto(license string, cli *mqtt.Client)
	DriveOut(cli *mqtt.Client)
}

type Parking struct {
	Id      int
	IsUsing bool
	License string
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
