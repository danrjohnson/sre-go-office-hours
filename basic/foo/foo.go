package foo

import "fmt"

type Foo struct {
	Name   string
	Weight int
	age    int
}

func (f *Foo) GetAge() int {
	return f.age
}

func (f *Foo) SetAge(age int) {
	f.age = age
}

func (f Foo) SetAge2(age int) {
	f.age = age
	fmt.Println(f.age)
}

func NewFoo(name string, weight int, age int) *Foo {
	return &Foo{
		Name:   name,
		Weight: weight,
		age:    age,
	}
}

func Whatever(f Foo) {
	fmt.Println(f.age)
}
