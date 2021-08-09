package zhihu_queue

import (
	"errors"
	"sync"
	"time"
)

// 一个订阅者就是一个管道，那么一个主题对应多个订阅者管道

type Broker interface {
	publish(topic string, msg interface{}) error        //消息推送
	subscribe(topic string) (<-chan interface{}, error) //主题订阅
	unsubscribe(topic string, sub <-chan interface{}) error
	close()
	broadcast(msg interface{}, subscribers []chan interface{}) // 消息广播，subscribers表示所有订阅者 每个请阅者是一个chan interface管道
	setConditions(capacity int)                                //用来设置条件，条件就是消息队列的容量
}

type BrokerImpl struct {
	exit     chan bool
	capacity int
	topics   map[string][]chan interface{} // key:主题 value:订阅者
	sync.RWMutex
}

func (b BrokerImpl) publish(topic string, msg interface{}) error {
	select {
		case <-b.exit:
			return errors.New("broker closed")
		default:
	}
	b.RLock()
	subscribers, ok := b.topics[topic]
	b.RUnlock()
	if !ok {
		return nil
	}
	b.broadcast(msg, subscribers)
	return nil
}

//订阅主题
func (b BrokerImpl) subscribe(topic string) (<-chan interface{}, error) {
	select {
	case <-b.exit:
		return nil, errors.New("broker closed")
	default:
	}
	ch := make(chan interface{}, b.capacity) //b.capacity表示管道的容量
	b.Lock()
	b.topics[topic] = append(b.topics[topic], ch)
	b.Unlock()
	return ch, nil
}

func (b BrokerImpl) unsubscribe(topic string, sub <-chan interface{}) error {
	select {
	case <-b.exit:
		return errors.New("broker closed")
	default:
	}
	b.RLock()
	subscribers, ok := b.topics[topic]
	b.RUnlock()
	if !ok {
		return nil
	}
	b.Lock()
	var newSubs []chan interface{}
	for _, subscriber := range subscribers {
		if subscriber == sub {
			continue
		}
		newSubs = append(newSubs, subscriber)
	}
	b.topics[topic] = newSubs
	b.Unlock()
	return nil
}

func (b BrokerImpl) close() {
	select {
	case <-b.exit:
		return
	default:
		close(b.exit)
		b.Lock()
		b.topics = make(map[string][]chan interface{})
		b.Unlock()
	}
}

func (b BrokerImpl) broadcast(msg interface{}, subscribers []chan interface{}) {
	count := len(subscribers)
	concurrency := 1
	switch {
	case count > 1000:
		concurrency = 3
	case count > 100:
		concurrency = 2
	default:
		concurrency = 1
	}
	pub := func(start int) {
		idleDuration := 5 * time.Millisecond
		idleTimeout := time.NewTimer(idleDuration)
		defer idleTimeout.Stop()
		for j := start; j < count; j += concurrency {
			if !idleTimeout.Stop() {
				select {
				case <-idleTimeout.C:
				default:
				}
			}
			idleTimeout.Reset(idleDuration)
			select {
			case subscribers[j] <- msg:
			case <-idleTimeout.C:
			case <-b.exit:
				return
			}
		}
	}
	for i := 0; i < concurrency; i++ {
		go pub(i)
	}
}

func (b BrokerImpl) setConditions(capacity int) {
	b.capacity = capacity
}

func NewBroker() *BrokerImpl {
	return &BrokerImpl{
		exit:   make(chan bool),
		topics: make(map[string][]chan interface{}),
	}
}
