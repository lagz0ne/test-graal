package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "graal_isolate_dynamic.h"
// #include "graal_isolate.h"
// #include "HelloWorld_dynamic.h"
// #include "HelloWorld.h"

import "C"

func main() {

	println("Hello world")
}
