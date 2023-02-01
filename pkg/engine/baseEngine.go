package engine

import "fmt"

const (
	CheckStatusStatus = "status"
)

type Engine interface {
	getEngineName() string
	getIterator() Iterator
}

type Iterator interface {
	hasNext() bool
	getNext() interface{}
}

type BaseEngine struct {
	Engine
}

func (e *BaseEngine) Run() {
	fmt.Printf("Running %s engine\n", e.getEngineName())

	iter := e.getIterator()

	for iter.hasNext() {
		fmt.Printf("%v,", iter.getNext())
	}
	fmt.Printf("\n")
}
