package tcp

import (
	"context"
	"example/project/godis/interface/tcp"
	"example/project/godis/logger"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Config struct {
	Address    string        `yaml:"address"`
	MaxConnect uint64        `yaml:"max-connect"`
	Timeout    time.Duration `yaml:"timeout"`
}


// ListenAndServeWithSignal 绑定端口并处理请求，直到收到停止信号为止
func ListenAndServeWithSignal(cfg *Config, handler tcp.Handler) error{
	closeChan := make(chan struct{})
	sigCh := make(chan os.Signal)
	// 告诉 signal ，将对应的信号通知 ch
	signal.Notify(sigCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT) //中断等信号
	go func() {
		select {
			case sig := <- sigCh:
				switch sig {
					case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
						closeChan <- struct{}{}
				}
		}
	}()
	//绑定监听的地址
	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	logger.Info(fmt.Sprintf("bind: %s, start listening...", cfg.Address))
	ListenAndServer(listener,handler,closeChan)
	return nil
}


// 服务端代码
func ListenAndServer(listener net.Listener, handler tcp.Handler, closeChan <-chan struct{}) {

	// 启动一个线程用来监控关闭信号 (优雅关闭)
	go func() {
		<-closeChan
		logger.Info("shutting down.....")
		listener.Close()
		handler.Close()
	}()

	//发生意外错误时关闭
	defer func() {
		listener.Close()
		handler.Close()
	}()

	ctx := context.Background()
	var wait sync.WaitGroup
	for {
		conn, err := listener.Accept() // Accept会一直阻塞直到有新的链接建立
		if err != nil {
			break
		}
		wait.Add(1)
		// 有新的连接请求就启动一个进程进行处理
		go func() {
			defer func() {
				wait.Done()
			}()
			handler.Handle(ctx,conn)
		}()
	}
	wait.Wait()

}
