package csdn

import (
	"fmt"
	"testing"
)

func TestIteratorPattern(t *testing.T) {
	namesRepository := new(NameRepository)
	for iter := namesRepository.GetIterator(); iter.HasNext(); {
		fmt.Println(iter.Next())
	}
}
