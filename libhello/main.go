package main

import "C"
import (
	"fmt"
)

const HELLO_WORLD = "hello world"

//export hello
func hello(in string) string {

	fmt.Println(in + " " + HELLO_WORLD)
	return HELLO_WORLD
}

func main() {

}