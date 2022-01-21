package cmap

import (
	"fmt"
	"testing"
)

func TestCmapNew(t *testing.T) {
	var concurrency int
	var pairRedistributor PairRedistributor
	cm, err := NewConcurrentMap(concurrency, pairRedistributor)
	if err == nil {
		t.Fatalf("No error when new a concurrent map with concurrency %d, but should not be the case!",
			concurrency)
	}
	concurrency = MAX_CONCURRENCY + 1
	cm, err = NewConcurrentMap(concurrency, pairRedistributor)
	if err == nil {
		t.Fatalf("No error when new a concurrent map with concurrency %d, but should not be the case!",
			concurrency)
	}
	concurrency = 16
	cm, err = NewConcurrentMap(concurrency, pairRedistributor)
	if err != nil {
		t.Fatalf("An error occurs when new a concurrent map: %s (concurrency: %d, pairRedistributor: %#v)",
			err, concurrency, pairRedistributor)
	}
	if cm == nil {
		t.Fatalf("Couldn't a new concurrent map! (concurrency: %d, pairRedistributor: %#v)",
			concurrency, pairRedistributor)
	}
	if cm.Concurrency() != concurrency {
		t.Fatalf("Inconsistent concurrency: expected: %d, actual: %d",
			concurrency, cm.Concurrency())
	}
}

func TestCmapPut(t *testing.T) {
	number := 30
	testCases := genTestingPairs(number)
	concurrency := 10
	var pairRedistributor PairRedistributor
	cm, _ := NewConcurrentMap(concurrency, pairRedistributor)
	var count uint64
	for _, p := range testCases {
		key := p.Key()
		element := p.Element()
		ok, err := cm.Put(key, element)
		if err != nil {
			t.Fatalf("An error occurs when putting a key-element to the cmap: %s (key: %s, element: %#v)", err, key, element)
		}
		if !ok {
			t.Fatalf("Couldn't put key-element to the cmap! (key: %s, element: %#v)", key, element)
		}
		actualElement := cm.Get(key)
		if actualElement == nil {
			t.Fatalf("Inconsistent element: expected: %#v, actual: %#v", element, nil)
		}
		ok, err = cm.Put(key, element)
		if err != nil {
			t.Fatalf("An error occurs when putting a repeated key-element to the cmap! %s (key: %s, element: %#v)", err, key, element)
		}
		if ok {
			t.Fatalf("Couldn't put key-element to the cmap! (key: %s, element: %#v)", key, element)
		}
		count++
		if cm.Len() != uint64(count) {
			t.Fatalf("Inconsistent size: expected: %d, actual: %d", count, cm.Len())
		}
	}
	if cm.Len() != uint64(number) {
		t.Fatalf("Inconsistent size: expected: %d, actual: %d", number, cm.Len())
	}
}

func TestCmapPutInParallel(t *testing.T) {
	number := 30
	testCases := genNoRepetitiveTestingPairs(number)
	concurrency := number / 2
	cm, _ := NewConcurrentMap(concurrency, nil)
	testingFunc := func(key string, element interface{}, t *testing.T) func(t *testing.T) {
		return func(t *testing.T) {
			t.Parallel()
			ok, err := cm.Put(key, element)
			if err != nil {
				t.Fatalf("An error occurs when putting a key-element to the cmap: %s (key: %s, element: %#v)",
					err, key, element)
			}
			if !ok {
				t.Fatalf("Couldn't put key-element to the cmap! (key: %s, element: %#v)",
					key, element)
			}
			actualElement := cm.Get(key)
			if actualElement == nil {
				t.Fatalf("Inconsistent element: expected: %#v, actual: %#v",
					element, nil)
			}
			ok, err = cm.Put(key, element)
			if err != nil {
				t.Fatalf("An error occurs when putting a repeated key-element to the cmap! %s (key: %s, element: %#v)",
					err, key, element)
			}
			if ok {
				t.Fatalf("Couldn't put key-element to the cmap! (key: %s, element: %#v)",
					key, element)
			}
		}
	}
	t.Run("ut in parallel", func(t *testing.T) {
		for _, p := range testCases {
			t.Run(fmt.Sprintf("Key=%s", p.Key()), testingFunc(p.Key(), p.Element(), t))
		}
	})
	if cm.Len() != uint64(number) {
		t.Fatalf("Inconsistent size: expected: %d, actual: %d",
			number, cm.Len())
	}
}
