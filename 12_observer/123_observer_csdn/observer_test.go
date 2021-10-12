package _23_observer_csdn

import (
	"fmt"
	"testing"
)

func TestObserverPattern(t *testing.T) {
	subject := new(Subject)
	new(HexaObserver).HexaObserver(subject)
	new(OctalObserver).OctalObserver(subject)
	new(BinaryObserver).BinaryObserver(subject)
	fmt.Println("First state change:15")
	subject.SetState(15)
	fmt.Println("Second state change:10")
	subject.SetState(10)
}
