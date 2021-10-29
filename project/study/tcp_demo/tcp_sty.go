package tcp_demo

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

// 用来测试一个地址的端口开放情况
func PortOpenTest() {
	hostname := flag.String("hostname", "", "hostname to test")
	startPort := flag.Int("startport", 80, "the port on which the scanning starts")
	endPort := flag.Int("endport", 100, "the port from which the scanning ends")
	timeout := flag.Duration("timeout", time.Second * 5, "timeout")
	flag.Parse()
	ports := make([]int, 0, 10)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go func(p int) {
			opened := isOpen(*hostname, port, *timeout)
			if opened {
				mutex.Lock()
				defer mutex.Unlock()
				ports = append(ports, p)
			}
		}(port)
	}
	wg.Wait()
	fmt.Println(ports)

}

func isOpen(host string, port int, timeout time.Duration) bool {
	time.Sleep(time.Millisecond * 1)
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), timeout)
	if err == nil {
		_ = conn.Close()
		return true
	}
	return false
}
