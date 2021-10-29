package loadbalance

import "errors"

type RoundRobinBalance struct {
	curIndex int
	allNodes []string
}

func (r *RoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	r.allNodes = append(r.allNodes, addr)
	return nil
}

func (r *RoundRobinBalance) Get() (string, error) {
	if len(r.allNodes) == 0 {
		return "", errors.New("list is empty")
	}
	lens := len(r.allNodes)
	if r.curIndex >= lens {
		r.curIndex = 0
	}
	curNode := r.allNodes[r.curIndex]
	r.curIndex = (r.curIndex + 1) % lens
	return curNode, nil
}
