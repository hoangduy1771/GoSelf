package main

import "fmt"
import "reflect"

var (
    name, course = "Duy", "Getting started with Go"
    module, clip = 4, 2
)

func main() {
    fmt.Println("Name and courses are set to", name, "and", course, ".")
    fmt.Println("Module and clip are set to", module, "and", clip, ".")
	fmt.Println("Name is of type: ", reflect.TypeOf(name))
	fmt.Println("Course is of type: ", reflect.TypeOf(module))
}