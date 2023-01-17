package pi

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
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

		// 检测逻辑部分
		if plateStr == "沪AZ0001" {
			wPipe.WriteString("pass")
			if cameraMode == FrontCamera {
				err = mgr.DriveIntoCar(plateStr)
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
				log.Printf("Error atoi data str: %s\n", err)
			}
			// log.Println(mgr.Spaces[id-1])
		} else {
			log.Printf("Error update data: %s\n", ErrArrayOutOfBounds)
		}
	}
}
