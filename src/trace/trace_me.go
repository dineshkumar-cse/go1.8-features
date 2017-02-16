package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

type Gopher struct {
	id   int
	name string
}

func (g Gopher) Say() {
	fmt.Println("Hi, I'm gopher:", g.id, "creator:", g.name)
}

func doSomething(name string, limit int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < limit; i++ {
		Gopher{i, name}.Say()
	}
}

func traceProgram() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	f, _ := os.Create("./dumps/trace.pprof")
	defer f.Close()
	defer trace.Stop()

	trace.Start(f)
	someLimit := 9000
	go doSomething("Hundred", someLimit, wg)
	go doSomething("Fifty", someLimit/3, wg)

	wg.Wait()
}

func main() {
	traceProgram()
}
