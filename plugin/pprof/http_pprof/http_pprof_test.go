package http_pprof

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"testing"
	"time"
)

// 随机生成字符串
func genSomeBytes() *bytes.Buffer {
	var buff bytes.Buffer
	for i := 0; i < 20000; i++ {
		buff.Write([]byte{'0' + byte(rand.Intn(10))})
	}
	return &buff
}

// 查看内存占用
func readMemStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf(" ===> Alloc:%d(bytes) HeapIdle:%d(bytes) HeapReleased:%d(bytes)", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
}

func test_mem() {
	container := make([]int, 8)
	log.Println("====> loop begin.")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
		if i == 16*1000*1000 {
			readMemStats()
		}
	}
	log.Println(" ====> loop end.")
}

func TestMemDemo(t *testing.T) {
	// 只要加这个，就能通过web方式来查看性能,记得引入包net/http/pprof
	// 访问 http://127.0.0.1:10000/debug/pprof/heap?debug=1
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:10000", nil))
	}()
	log.Println(" ===> [Start].")
	readMemStats()
	test_mem()
	readMemStats()
	log.Println(" ===> [force gc].")
	runtime.GC() //强制调用gc回收
	log.Println(" ===> [Done].")
	readMemStats()
	time.Sleep(3600 * time.Second) //睡眠，保持程序不退出
}

func TestCPUDemo(t *testing.T) {
	go func() {
		for {
			log.Println(" ===> loop begin.")
			for i := 0; i < 1000; i++ {
				genSomeBytes()
			}
			log.Println(" ===> loop end.")
		}
	}()
	// http://127.0.0.1:10000/debug/pprof
	http.ListenAndServe("0.0.0.0:10000", nil)
}
