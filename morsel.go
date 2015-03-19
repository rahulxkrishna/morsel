package main

import (
//"time"
)

func main() {
	//TODO spawn of goroutines to handle each operation

	var v MView

	m := &Model{}
	c := &Controller{}
	v = &CLView{}

	v.Init(m, c)
	c.Init(m, v)

	m.run()
	v.Run()
}
