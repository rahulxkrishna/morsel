package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const MAX_FEEDS = 24

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Desc  string `xml:"description"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title    string `xml:"title"`
	Desc     string `xml:"description"`
	Link     string `xml:"link"`
	Category string `xml:"category"`
}

func readConf() ([]string, error) {
	file, err := os.Open("feedme.conf")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rssList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rssList = append(rssList, scanner.Text())
	}

	return rssList, scanner.Err()
}

func fetchRSS() ([]Item, error) {
	var rssText []Item
	rssList, _ := readConf()

	numFeeds := len(rssList)
	maxPerFeed := MAX_FEEDS / numFeeds
	feedCount := 1

	for _, src := range rssList {
		response, err := http.Get(src)
		if err != nil {
			return rssText, err
		}

		xdat, err := ioutil.ReadAll(response.Body)
		var rss RSS
		xml.Unmarshal(xdat, &rss)

		for i, item := range rss.Channel.Items {
			title := fmt.Sprintf("%d) %s\n", feedCount, item.Title)
			desc := fmt.Sprintf("%s\n", item.Desc)
			item := Item{title, desc, "", ""}
			rssText = append(rssText, item)
			if i > maxPerFeed {
				break
			}
			feedCount++
		}
	}

	return rssText, nil
}

func main() {
	for {
		rssList, _ := fetchRSS()
		displayRSS(rssList)
		time.Sleep(5 * time.Minute)
	}
}
