package main

import "github.com/dmitsh/go-stuff/pkg/engine"

type Worker interface {
	Run()
}

func main() {
	var worker Worker = engine.NewIntEngine()
	worker.Run()

	worker = engine.NewCharEngine("qwertyuiop")
	worker.Run()
}
