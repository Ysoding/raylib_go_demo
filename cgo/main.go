package main

/*
#cgo LDFLAGS: -L. -lexample
#include "example.h"
*/
import "C"
import "fmt"

func main() {
	a := C.int(3)
	b := C.int(4)
	result := C.add(a, b)
	fmt.Printf("Result: %d\n", result)
}
