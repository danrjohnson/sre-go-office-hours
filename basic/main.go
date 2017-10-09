package main

import "fmt"
import "github.com/danrjohnson/office-hours/basic/foo"

func main() {
	fmt.Println("Hello everyone.")

	theVariable := foo.NewFoo("Dan", 180, 32)
	fmt.Println(theVariable)
	fmt.Printf("My age is %d\n", theVariable.GetAge())
	theVariable.SetAge(29)
	fmt.Printf("My age is %d\n", theVariable.GetAge())
	theVariable.SetAge2(70)
	fmt.Printf("My age is %d\n", theVariable.GetAge())
	foo.Whatever(*theVariable)
}
