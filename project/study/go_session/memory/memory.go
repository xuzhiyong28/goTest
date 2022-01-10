package memory

import (
	"container/list"
	"example/project/study/go_session"
	"sync"
	"time"
)

type SessionStore struct {
	sid          string                      // session id唯一标示
	timeAccessed time.Time                   // 最后访问时间
	value        map[interface{}]interface{} // session里面存储的值
}

func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
}

func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sid
}

type MemoryProvider struct {
	lock     sync.Mutex               // 用来锁
	sessions map[string]*list.Element // 用来存储在内存
	list     *list.List               // 用来做 gc
}


func (pder *MemoryProvider) SessionInit(sid string) (go_session.Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newSession := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushFront(newSession)
	pder.sessions[sid] = element
	return newSession, nil
}

func (pder *MemoryProvider) SessionRead(sid string) (go_session.Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (pder *MemoryProvider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
	}
	return nil
}

func (pder *MemoryProvider) SessionGC(maxLifeTime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxLifeTime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (pder *MemoryProvider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}

var pder = &MemoryProvider{
	list: list.New(),
	sessions: make(map[string]*list.Element, 0),
}

func init() {
	go_session.Register("memory", pder)
}
