package pkg_errors

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

type ErrorTyped struct {
	error
}

func WrappedError(e error) error {
	return errors.Wrap(e, "An error occurred in WrappedError")
}

func TestDemo0(t *testing.T) {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")
	fmt.Printf("%+v\n", err)
}

func TestDemo1(t *testing.T) {
	e := errors.New("standard error")
	fmt.Println("Regular Error - ", WrappedError(e))
	fmt.Println("Typed Error - ", WrappedError(ErrorTyped{errors.New("typed error")}))
	fmt.Println("Nil -", WrappedError(nil))
}

func TestUnWrap(t *testing.T) {
	err := error(ErrorTyped{errors.New("an error occurred")})
	err = errors.Wrap(err, "wrapped")
	fmt.Println("wrapped error: ", err)
	switch errors.Cause(err).(type) {
	case ErrorTyped:
		fmt.Println("a typed error occurred: ", err)
	default:
		fmt.Println("an unknown error occurred")
	}
}
