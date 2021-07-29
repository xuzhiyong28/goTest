package multiple

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestServiceDemo1(t *testing.T) {
	ServiceDemo1()
}

func TestClientDemo1(t *testing.T) {
	ClientDemo1()
}


func TestOther(t *testing.T) {
	var challenge [32]byte
	rand.Read(challenge[:])
	fmt.Println(challenge)
}