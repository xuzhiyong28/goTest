package service

import (
	"errors"
	"fmt"
	"sync/atomic"
)

var (
	// ErrAlreadyStarted is returned when somebody tries to start an already
	// running service.
	ErrAlreadyStarted = errors.New("already started")
	// ErrAlreadyStopped is returned when somebody tries to stop an already
	// stopped service (without resetting it).
	ErrAlreadyStopped = errors.New("already stopped")
	// ErrNotStarted is returned when somebody tries to stop a not running
	// service.
	ErrNotStarted = errors.New("not started")
)

type Service interface {
	Start() error
	OnStart() error
	Stop() error
	OnStop()
	Reset() error
	OnReset() error
	IsRunning() bool
	Quit() <-chan struct{}
	String() string
	Wait()
}

type BaseService struct {
	name    string
	started uint32 // atomic
	stopped uint32 // atomic
	quit    chan struct{}
	impl    Service
}

func NewBaseService(name string, impl Service) *BaseService {
	return &BaseService{
		name: name,
		quit: make(chan struct{}),
		impl: impl,
	}
}

func (bs *BaseService) Start() error {
	if atomic.CompareAndSwapUint32(&bs.started, 0, 1) {
		if atomic.LoadUint32(&bs.stopped) == 1 {
			// error
			atomic.StoreUint32(&bs.started, 0)
			return ErrAlreadyStopped
		}

		if err := bs.impl.OnStart(); err != nil {
			atomic.StoreUint32(&bs.started, 0)
			return err
		}
		return nil
	}
	return ErrAlreadyStarted
}

func (bs *BaseService) OnStart() error { return nil }

func (bs *BaseService) Stop() error {
	if atomic.CompareAndSwapUint32(&bs.stopped, 0, 1) {
		if atomic.LoadUint32(&bs.started) == 0 {
			atomic.StoreUint32(&bs.stopped, 0)
			return ErrNotStarted
		}
		bs.impl.OnStop()
		close(bs.quit)
		return nil
	}
	return ErrAlreadyStopped
}

func (bs *BaseService) OnStop() {}

func (bs *BaseService) Reset() error {
	if !atomic.CompareAndSwapUint32(&bs.stopped, 1, 0) {
		return fmt.Errorf("can't reset running %s", bs.name)
	}
	atomic.CompareAndSwapUint32(&bs.started, 1, 0)
	bs.quit = make(chan struct{})
	return bs.impl.OnReset()
}

func (bs *BaseService) OnReset() error {
	panic("The service cannot be reset")
}

func (bs *BaseService) IsRunning() bool {
	return atomic.LoadUint32(&bs.started) == 1 && atomic.LoadUint32(&bs.stopped) == 0
}

func (bs *BaseService) Wait() {
	<-bs.quit
}

func (bs *BaseService) String() string {
	return bs.name
}

func (bs *BaseService) Quit() <-chan struct{} {
	return bs.quit
}
