package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//setup test vars
	var firstVar func() = firstFunc
	var secondVar func() = secondFunc

	//ideally these bytes will be shellcode/module code
	var testBytes []byte = []byte{40, 96, 0, 0, 192, 0, 0, 0}

	//get pointer to bytes and functions
	testPtr := &testBytes
	firstPrt := &firstVar
	secondPtr := &secondVar

	fmt.Println(" SecondPtr:", secondPtr)

	//create variable which stores function pointer
	var res func()

	//modify function pointer to demo proof of concept
	res = *firstPrt
	res()
	res = *secondPtr
	res()

	//now the goal here is to write the bytecode for an arbitrary function
	//into memory, get a pointer to that memory, then set our function pointer
	//res variable to point to that memory. The hypothesis is that once we call
	//the res variable, we should get our function to execute. We may need the
	//go/plan9 asm equivalent of jmp instructions to get to our bytecode.
	conv := (*func())(unsafe.Pointer(testPtr))
	fmt.Println("conv:", conv, "&conv:", &conv, "*conv", *conv)
	res = *conv

	fmt.Printf(" conv %v\n &conv %v\n testPtr %v\n &testPtr %v\n  res %v\n &res %v\n testBytes %s\n &testBytes %v\n", conv, &conv, *testPtr, &testPtr, res, &res, testBytes, &testBytes)
	res()
}

func firstFunc() {
	fmt.Println("First Function")
}

func secondFunc() {
	fmt.Println("Second Function")
}
