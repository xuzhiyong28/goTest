package loadbalance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
	curIdx   int
	allNodes []string
}

func (r *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	r.allNodes = append(r.allNodes, addr)
	return nil
}

func (r *RandomBalance) Get() (string, error) {
	if len(r.allNodes) == 0 {
		return "", errors.New("allNodes is empty")
	}
	r.curIdx = rand.Intn(len(r.allNodes))
	return r.allNodes[r.curIdx], nil
}
