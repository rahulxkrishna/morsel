package main

import (
	"bufio"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const MAX_FEEDS = 20

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

// readConf reads the feeds list from the 'morsel.conf' file
func readConf() ([]string, error) {
	file, err := os.Open("morsel.conf")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rssList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if (strings.Trim(line, " "))[0] == '#' {
			continue
		}
		rssList = append(rssList, line)
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
