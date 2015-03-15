package main

import (
//"time"
)

func main() {

	//TODO spawn of goroutines to handle each operation

	m := &Model{}
	c := &Controller{}
	v := &View{}

	c.Init(m, v)
	v.Init(m, c)

	m.run()
	v.run()
}
