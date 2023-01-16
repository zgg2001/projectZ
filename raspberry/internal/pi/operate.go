package pi

import (
	"bufio"
	"log"
	"os"
	"os/exec"
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
func PythonRunTask() {

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
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Panic("Error read bytes: ", err)
		}
		plate_str := strings.TrimRight(string(line), "\n")
		log.Println(plate_str)

		// 检测逻辑部分
		if plate_str == "沪AZ0001" {
			wPipe.WriteString("pass")
		} else {
			wPipe.WriteString("reject")
		}
	}
}
