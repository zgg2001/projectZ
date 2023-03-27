package operate

import (
	"strconv"

	"github.com/zgg2001/projectZ/server/parking_server/internal/data"
	"github.com/zgg2001/projectZ/server/parking_server/pkg/rpc"
)

var ServerService = &serverService{}

type ServiceOperation interface {
	DataInit()
}

type serverService struct {
	rpc.UnimplementedProjectServiceServer
	funcChan chan func()
}

func (ss *serverService) DataInit() error {
	ss.funcChan = make(chan func(), 22)
	err := data.UserInit()
	if err != nil {
		return err
	}
	err = data.ParkingInit()
	if err != nil {
		return err
	}
	return nil
}

func (ss *serverService) DBMgrTaskQueueRunning() {
	for {
		f, ok := <-ss.funcChan
		if !ok {
			break
		}
		f()
	}
}

func (ss *serverService) SetParkingSpace(license string, pid, sid int32, requestTime int64, mode int32) {
	// redis
	row := &data.RecordRow{
		License:   license,
		PId:       pid,
		SId:       sid,
		EntryTime: requestTime,
	}
	data.RedisAddRecord(row)
	// mysql
	ss.funcChan <- func() {
		if mode == RearCamera {
			data.DeleteRecordTbl(license)
		} else {
			data.InsertRecordTbl(license, pid, sid, requestTime)
		}
		data.InsertParkingRecordTbl(license, pid, sid, mode, requestTime)
	}
}

func (ss *serverService) SetBalance(uid string, balance int32) {
	// redis
	data.RedisSetBalance(uid, balance)
	// mysql
	ss.funcChan <- func() {
		data.ChangeUserBalanceTbl(uid, balance)
	}
}

func (ss *serverService) UpdateParkingData(pid, temperature, humidity, weather int32) {
	// redis
	data.RedisSetParkingInfo(pid, temperature, humidity, weather)
}

func (ss *serverService) UpdateParkingSpaceData(pid, sid, temperature, humidity, alarm int32) {
	// redis
	value := strconv.Itoa(int(temperature)) + "+" + strconv.Itoa(int(humidity)) + "+" + strconv.Itoa(int(alarm))
	data.RedisSetParkingSpaceInfo(pid, sid, value)
}
