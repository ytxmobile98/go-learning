package main

import "fmt"

type A struct {
	x int
}

type AB struct {
	*A
	y int
}

type Foo struct {
	*AB
}

func main() {
	foo := Foo{&AB{&A{1}, 2}}
	fmt.Println(foo)
	fmt.Println(foo.x)
}
