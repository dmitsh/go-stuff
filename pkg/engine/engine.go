package engine

type IntEngine struct {
	BaseEngine

	start  int
	length int
}

func NewIntEngine() *IntEngine {
	e := &IntEngine{
		start:  0,
		length: 10,
	}
	e.engine = e
	return e
}

func (e *IntEngine) getEngineName() string {
	return "IntEngine"
}

func (e *IntEngine) getIterator() Iterator {
	return &IntIterator{
		start:  e.start,
		length: e.length,
	}
}

type IntIterator struct {
	start  int
	length int
	index  int
}

func (i *IntIterator) hasNext() bool {
	return i.index < i.length
}

func (i *IntIterator) getNext() interface{} {
	ret := i.start + i.index
	i.index++
	return ret
}
