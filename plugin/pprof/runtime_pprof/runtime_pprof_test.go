package runtime_pprof

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"testing"
	"time"
)

//一段有问题的程序
func logicErrorCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan, value:%v\n", v)
		default:
		}
	}
}

func createPathFile(fileName string) string {
	path, _ := os.Getwd()
	var ostype = runtime.GOOS
	filePath := ""
	if ostype == "windows" {
		filePath = path + "\\" + fileName
	} else if ostype == "linux" {
		filePath = path + "/" + fileName
	}
	return filePath
}


func TestPprofCpu(t *testing.T) {
	file, _ := os.Create(createPathFile("cpu.pprof"))
	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()
	for i := 0; i < 8; i++ {
		go logicErrorCode()
	}
	time.Sleep(20 * time.Second)
}

func TestPprofMem(t *testing.T) {
	file, _ := os.Create(createPathFile("cpu.pprof"))
	for i := 0; i < 8; i++ {
		go logicErrorCode()
	}
	time.Sleep(20 * time.Second)
	pprof.WriteHeapProfile(file)
	file.Close()
}
