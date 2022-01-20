package cmap

// 接口
type ConcurrentMap interface {
	// Concurrency 会返回并发量。
	Concurrency() int
	// Put 会推送一个键-元素对。
	// 注意！参数element的值不能为nil。
	// 第一个返回值表示是否新增了键-元素对。
	// 若键已存在，新元素值会替换旧的元素值。
	Put(key string, element interface{}) (bool, error)
	// Get 会获取与指定键关联的那个元素。
	// 若返回nil，则说明指定的键不存在。
	Get(key string) interface{}
	// Delete 会删除指定的键-元素对。
	// 若结果值为true则说明键已存在且已删除，否则说明键不存在。
	Delete(key string) bool
	// Len 会返回当前字典中键-元素对的数量。
	Len() uint64
}

//实现
type myConcurrentMap struct {
	concurrency int
	// 分段数组
	segments []Segment
	total    uint64
}

func (cmap *myConcurrentMap) Concurrency() int {
	panic("implement me")
}

func (cmap *myConcurrentMap) Put(key string, element interface{}) (bool, error) {
	panic("implement me")
}

func (cmap *myConcurrentMap) Get(key string) interface{} {
	panic("implement me")
}

func (cmap *myConcurrentMap) Delete(key string) bool {
	panic("implement me")
}

func (cmap *myConcurrentMap) Len() uint64 {
	panic("implement me")
}

func NewConcurrentMap(concurrency int, pairRedistributor PairRedistributor) (ConcurrentMap, error) {
	if concurrency <= 0 {
		return nil, newIllegalParameterError("concurrency is too small")
	}
	if concurrency > MAX_CONCURRENCY {
		return nil, newIllegalParameterError("concurrency is too large")
	}
	cmap := &myConcurrentMap{}
	cmap.concurrency = concurrency
	cmap.segments = make([]Segment, concurrency)
	for i := 0; i < concurrency; i++ {
		cmap.segments[i] = newSegment(DEFAULT_BUCKET_NUMBER, pairRedistributor)
	}
	return cmap, nil
}
