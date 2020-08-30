package tools

import (
	"fmt"
	"runtime"
)

func trace2() {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
}

func executeFirst() {
	trace2()
}

func executeLast() {
	trace2()
}

// DeferFunctionCall shows the execution order of a defered function
func DeferFunctionCall() {
	executeFirst()
	defer fmt.Println("world")
	fmt.Println("hello")
	executeLast()
}
