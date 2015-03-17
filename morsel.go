package main

import (
//"time"
)

func main() {
	//TODO spawn of goroutines to handle each operation

	m := &Model{}
	c := &Controller{}
	v := &View{}

	v.Init(m, c)
	c.Init(m, v)

	m.run()
	v.run()
}
