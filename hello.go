package main

import (
	"fmt"
	"reflect"
)

func main() {
	var version float32 = 1.1 //full declaration
	var name = "Tomas"        //type inference
	age := 24                 //short declaration
	fmt.Println("Hello mr.", name, " your age is ", age)
	fmt.Println("This program is on version ", version)
	fmt.Println("The type of name is : ", reflect.TypeOf(name))
}
