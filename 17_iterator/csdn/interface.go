package csdn


type Container interface {
	GetIterator() Iterator
}

type SomeSlice interface{}

type Iterator interface {
	HasNext() bool
	Next() SomeSlice
}


