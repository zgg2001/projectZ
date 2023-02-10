package pi

import (
	"bufio"
	"context"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/zgg2001/projectZ/raspberry/internal/transmission"
	"github.com/zgg2001/projectZ/server/pkg/rpc"
)

func PipeInit() {

	if _, err := os.Stat(WordDir); err != nil {
		err := os.MkdirAll(WordDir, 0711)
		if err != nil {
			log.Panic("Error creating directory: ", err)
			return
		}
	}

	os.Remove(PipeRead)
	os.Remove(PipeWrite)
	err := syscall.Mkfifo(PipeRead, 0666)
	if err != nil {
		log.Panic("Error make read fifo", err)
	}
	err = syscall.Mkfifo(PipeWrite, 0666)
	if err != nil {
		log.Panic("Error make write fifo", err)
	}

	log.Println("Pipeline init complete")
}

func PipeRemove() {
	os.Remove(PipeRead)
	os.Remove(PipeWrite)
	log.Println("Pipeline remove complete")
}

// 硬件脚本
func PythonStartUp() *exec.Cmd {
	cmd := exec.Command(PythonPath, ScriptPath)
	err := cmd.Start()
	if err != nil {
		log.Panic("Error start python script", err)
	}
	log.Println("Python script startup")
	return cmd
}

func PythonCancel(cmd *exec.Cmd) {
	cmd.Process.Signal(syscall.SIGQUIT)
	log.Println("Python cancel complete")
}

// 树莓派硬件数据交互
func RunPythonTask(mgr *ParkingMgr) {

	log.Println("Run python task ... ")

	rPipe, err := os.OpenFile(PipeRead, os.O_RDWR, os.ModeNamedPipe)
	if err != nil {
		log.Panic("Error open file: ", err)
	}
	wPipe, err := os.OpenFile(PipeWrite, os.O_RDWR, 0777)
	if err != nil {
		log.Panic("Error open file: ", err)
	}

	conn := transmission.RPCNewClient()
	defer conn.Close()
	rpcClient := rpc.NewProjectServiceClient(conn)

	reader := bufio.NewReader(rPipe)
	for {
		// parse
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Panic("Error read bytes: ", err)
		}
		strArr := strings.Split(string(line), ":")
		if len(strArr) < 2 {
			log.Printf("Error split str: %s\n", ErrProtocol)
			continue
		}
		cameraMode, err := strconv.Atoi(strArr[0])
		if err != nil {
			log.Printf("Error atoi str: %s\n", err)
			continue
		}
		plateStr := strings.TrimRight(strArr[1], "\n")
		log.Println(cameraMode, plateStr)
		sid, err := mgr.FindEmptySpace()
		if err != nil {
			log.Println("Error find empty space: ", err)
			continue
		}

		request := &rpc.LPCheckRequest{
			Model:          int32(cameraMode),
			ParkingId:      ParkingID,
			ParkingSpaceId: int32(sid + 1),
			License:        plateStr,
		}
		resp, err := rpcClient.LicencePlateCheck(context.Background(), request)
		if err != nil {
			log.Printf("Error licence plate check: %s\n", err)
			wPipe.WriteString("reject")
			continue
		}
		log.Println("Check success", resp.Result)

		// 检测逻辑部分
		if resp.Result == transmission.LPCheckSucceeded {
			wPipe.WriteString("pass")
			if cameraMode == FrontCamera {
				err = mgr.DriveIntoCar(plateStr, sid)
			} else {
				err = mgr.DriveOutCar(plateStr)
			}
			if err != nil {
				log.Println("Error operator", err)
			}
		} else {
			wPipe.WriteString("reject")
		}
	}
}

// esp32数据收集
func RunDataTask(dataChan chan string, mgr *ParkingMgr) {

	log.Println("Run data task ... ")

	for {

		dataStr, ok := <-dataChan
		if !ok {
			break
		}

		strArr := strings.Split(dataStr, ":")
		if len(strArr) < 5 {
			log.Printf("Error split data str: %s\n", ErrProtocol)
			continue
		}
		id, err := strconv.Atoi(strArr[0])
		if err != nil {
			log.Printf("Error atoi data id: %s\n", err)
			continue
		}

		// update
		if len(mgr.Spaces) >= id {
			err = mgr.Spaces[id-1].UpdataData(strArr)
			if err != nil {
				log.Printf("Error updata data str: %s\n", err)
			}
			// log.Println(mgr.Spaces[id-1])
		} else {
			log.Printf("Error update data: %s\n", ErrArrayOutOfBounds)
		}
	}
}

// 停车场数据上传
func UploadPiData(mgr *ParkingMgr) {

	log.Println("Upload Pi data ... ")
	time.Sleep(time.Second * 5)

	conn := transmission.RPCNewClient()
	defer conn.Close()
	rpcClient := rpc.NewProjectServiceClient(conn)

	parkingSpaceCount := len(mgr.Spaces)
	for parkingSpaceCount <= 0 {
		parkingSpaceCount = len(mgr.Spaces)
		time.Sleep(time.Second)
	}

	packet := &rpc.UploadInfoRequest{}
	packet.PInfo = &rpc.ParkingInfo{
		Id:          ParkingID,
		Temperature: 4,
		Humidity:    28,
		Weather:     WeatherSunny,
	}
	for id := 0; id < parkingSpaceCount; id++ {
		info := rpc.ParkingSpaceInfo{
			Id:          int32(id) + 1,
			Temperature: 0,
			Humidity:    0,
			Alarm:       NoAlarm,
		}
		packet.InfoArr = append(packet.InfoArr, &info)
	}

	for {
		// get and upload
		for id := 0; id < parkingSpaceCount; id++ {
			temperature, humidity, alarm := mgr.Spaces[id].GetData() // real id = id + 1
			packet.InfoArr[id].Temperature = temperature
			packet.InfoArr[id].Humidity = humidity
			packet.InfoArr[id].Alarm = alarm
		}
		_, err := rpcClient.UploadParkingInfo(context.Background(), packet)
		if err != nil {
			log.Printf("Error upload parking info: %s\n", err)
		}
		time.Sleep(time.Second * 5)
	}
}
