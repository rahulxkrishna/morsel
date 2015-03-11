package main

import (
	"bufio"
	"encoding/xml"
	_ "fmt"
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

// readConf reads the feeds list from the 'feedme.conf' file
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

func fetchRSS() ([]Channel, error) {
	var feeds []Channel
	rssList, _ := readConf()

	for _, src := range rssList {
		response, err := http.Get(src)
		if err != nil {
			return feeds, err
		}

		xdat, err := ioutil.ReadAll(response.Body)
		var rss RSS
		xml.Unmarshal(xdat, &rss)

		feeds = append(feeds, rss.Channel)

	}

	return feeds, nil
}

func main() {

	for {
		feeds, _ := fetchRSS()
		displayRSS(feeds)
		time.Sleep(5 * time.Minute)
	}
}
