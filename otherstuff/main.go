package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func Foo() (int, string, bool) {
	return 42, "hello", true
}

func OpenThing() (int, error) {
	_, err := os.OpenFile("/tmp/does/not/exist/lol/peeps", os.O_RDONLY, os.FileMode(0600))
	if err != nil {
		return 0, err
	}
	_, err = os.OpenFile("/tmp/does/not/exist/", os.O_RDONLY, os.FileMode(0600))
	if err != nil {
		return 0, nil
	}
	return 0, nil
}

func ErrorThrower() error {
	return errors.New("Hi I'm An Error!")
}

func main() {
	foo, bar, baz := Foo()
	fmt.Println(foo, bar, baz)
	err := ErrorThrower()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("howdy!")
}
