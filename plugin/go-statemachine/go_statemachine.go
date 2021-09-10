package go_statemachine

import (
	"fmt"
	"github.com/filecoin-project/go-statemachine"
)

var _ statemachine.StateHandler = &TestHander{}



type TestHander struct {
	proceed chan struct{}
	done    chan struct{}
}

func (t *TestHander) Plan(events []statemachine.Event, state interface{}) (interface{}, uint64, error) {
	return t.plan(events, state.(*statemachine.TestState))
}

func (t *TestHander) plan(events []statemachine.Event, state *statemachine.TestState) (func(ctx statemachine.Context, myState statemachine.TestState) error, uint64, error) {
	for _, event := range events {
		e := event.User.(*statemachine.TestEvent)
		switch e.A {
		case "start":
			state.A = 1
		case "b":
			state.A = 2
			state.B = e.Val
		}
	}

	switch state.A {
	case 1:
		return t.step0, uint64(len(events)), nil
	case 2:
		return t.step1, uint64(len(events)), nil
	default:
		fmt.Println("========error========")
	}
	panic("how?")
}

func (t *TestHander) step0(ctx statemachine.Context, st statemachine.TestState) error {
	fmt.Println("=== step0 ===")
	ctx.Send(&statemachine.TestEvent{A: "b", Val: 55})
	<-t.proceed
	return nil
}

func (t *TestHander) step1(ctx statemachine.Context, st statemachine.TestState) error {
	fmt.Println("=== step1 ===")
	close(t.done)
	return nil
}
