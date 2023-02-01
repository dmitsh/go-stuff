package engine

type CharEngine struct {
	BaseEngine

	str string
}

func NewCharEngine(str string) *CharEngine {
	e := &CharEngine{
		str: str,
	}
	e.engine = e
	return e
}

func (e *CharEngine) getEngineName() string {
	return "CharEngine"
}

func (e *CharEngine) getIterator() Iterator {
	return &CharIterator{
		str: e.str,
	}
}

type CharIterator struct {
	str   string
	index int
}

func (i *CharIterator) hasNext() bool {
	return i.index < len(i.str)
}

func (i *CharIterator) getNext() interface{} {
	ret := string(i.str[i.index])
	i.index++
	return ret
}
