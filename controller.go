package main

import (
//"fmt"
)

var curPos int
var feeds []Feed

func getNextFeeds() {
	n := 5
	id := 1

	for _, morsel := range morsels {
		if curPos >= len(morsel.Items) {
			curPos = 0
		}

		if curPos+n > len(morsel.Items) {
			n = len(morsel.Items) - curPos
		}

		for j := 0; j < n; j++ {
			feeds = append(feeds, Feed{id, morsel.Title, morsel.Items[curPos+j].Title, ""})
			id += 1
		}
	}
	curPos += n
	displayFeeds(feeds)
}

func getPrevFeeds() {
	//displayFeeds(morsels)
}

func handleInput(ip string) {
	switch ip {
	case "n":
	case "":
		getNextFeeds()
	case "p":
		getPrevFeeds()
	}
}
