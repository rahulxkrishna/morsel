package main

import (
	"time"
)

func main() {

	for {
		feeds, _ := fetchRSS()
		displayRSS(feeds)
		time.Sleep(5 * time.Minute)
	}
}
