package v1

/**
假设需要创建一个文件存放数据，同一个时刻可能会有多个goroutine分别对该文件进行写操作和读操作。
每一次写操作都向该文件写入若干个字符数据，这若干字符的数据作为一个独立的数据块，多个写操作之间不能导致数据块的穿插混淆。
每一次读操作都能从文件读取一个数据块，且多个goruntine需要按顺序读取。
*/
import (
	"errors"
	"io"
	"os"
	"sync"
)

// Data 代表数据的类型
type Data []byte

// DataFile 代表数据文件的接口类型
type DataFile interface {
	// Read 会读取一个数据块。
	Read() (rsn int64, d Data, err error)
	// Write 会写入一个数据块。
	Write(d Data) (wsn int64, err error)
	// RSN 会获取最后读取的数据块的序列号。
	RSN() int64
	// WSN 会获取最后写入的数据块的序列号。
	WSN() int64
	// DataLen 会获取数据块的长度。
	DataLen() uint32
	// Close 会关闭数据文件。
	Close() error
}

// myDataFile 代表数据文件的实现类型。
type myDataFile struct {
	f       *os.File     // 文件
	fmutex  sync.RWMutex // 被用于文件的读写锁
	woffset int64        // 写操作需要用到的偏移量。
	roffset int64        // 读操作需要用到的偏移量。
	wmutex  sync.Mutex   // 写woffset需要用到的互斥锁。
	rmutex  sync.Mutex   // 读roffset需要用到的互斥锁。
	dataLen uint32       // 数据块长度。
}

// NewDataFile 会新建一个数据文件的实例。
func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}
	df := &myDataFile{f: f, dataLen: dataLen}
	return df, nil
}


func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	// 读取并更新写偏移量
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset = df.woffset + int64(df.dataLen)
	df.wmutex.Unlock()

	//写入一个数据块。
	wsn = offset / int64(df.dataLen)
	var bytes []byte
	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen] //如果data超过设定的值，就只取前dataLen个
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	return wsn, err
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset = df.roffset + int64(df.dataLen)
	df.rmutex.Unlock()
	rsn = offset / int64(df.dataLen)

	bytes := make([]byte, df.dataLen)
	for {
		df.fmutex.RLock()
		_, err := df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				df.fmutex.RUnlock()
				continue
			}
			df.fmutex.RUnlock()
			return
		} else {
			d = bytes
			df.fmutex.RUnlock()
		}
		return
	}

	return rsn, nil, nil
}

func (df *myDataFile) RSN() int64 {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) WSN() int64 {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

func (df *myDataFile) Close() error {
	if df.f == nil {
		return nil
	}
	return df.f.Close()
}
