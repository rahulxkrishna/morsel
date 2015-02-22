package main

import (
	"encoding/xml"
	"fmt"
	//"github.com/buger/goterm"
	"io/ioutil"
	"net/http"
	_ "os"
	"time"
)

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Category    string `xml:"category"`
}

func fetchRSS() {
	response, err := http.Get("http://feeds.bbci.co.uk/news/world/rss.xml")
	if err != nil {
		return
	}

	xdat, err := ioutil.ReadAll(response.Body)
	var rss RSS
	xml.Unmarshal(xdat, &rss)

	for i, item := range rss.Channel.Items {
		fmt.Printf("%d) %s\n", i+1, item.Title)
		fmt.Printf("\t. %s\n", item.Description)
		//fmt.Printf("\t. %s\n", item.Link)

		if i > 10 {
			break
		}
	}
}

func main() {
	for {
		fetchRSS()
		time.Sleep(5000 * time.Millisecond)
	}
}
