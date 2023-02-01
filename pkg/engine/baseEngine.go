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

// Working
type BaseEngine struct {
	engine Engine
}

func (e *BaseEngine) Run() {
	fmt.Printf("Running %s engine\n", e.engine.getEngineName())

	iter := e.engine.getIterator()

	for iter.hasNext() {
		fmt.Printf("%v,", iter.getNext())
	}
	fmt.Printf("\n")
}

// Not working

/*
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


*/
