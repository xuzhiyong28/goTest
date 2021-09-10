package go_statemachine

import (
	"fmt"
	"github.com/filecoin-project/go-statemachine"
	"github.com/ipfs/go-datastore"
	logging "github.com/ipfs/go-log/v2"
	"testing"
)

func init() {
	logging.SetLogLevel("*", "INFO") // nolint: errcheck
}



func TestBasic(t *testing.T) {
	ds := datastore.NewMapDatastore()
	th := &TestHander{
		done:    make(chan struct{}),
		proceed: make(chan struct{}),
	}
	close(th.proceed)
	smm := statemachine.New(ds, th, statemachine.TestState{})
	err := smm.Send(uint64(2), &statemachine.TestEvent{A : "start"})
	if err != nil {
		fmt.Print(err)
	}
	<-th.done
	fmt.Println("=== end ===")
}
