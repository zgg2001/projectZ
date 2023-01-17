package pi

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
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
	// 硬件逻辑
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
func PythonRunTask(cli *mqtt.Client, count int) {

	var mgr ParkingMgr
	mgr.Init(count, cli)

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
		}
		cameraMode, err := strconv.Atoi(strArr[0])
		if err != nil {
			log.Printf("Error split str: %s\n", ErrProtocol)
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
