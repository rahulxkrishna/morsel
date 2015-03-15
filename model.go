package main

import (
	"bufio"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

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

type Feed struct {
	Id     int
	Source string
	Title  string
	Desc   string
	Link   string
}

type Model struct {
	feeds []Feed
}

// readConf reads the feeds list from the 'morsel.conf' file
func (m *Model) readConf() ([]string, error) {
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

func (m *Model) getFeed(id int) (Feed, error) {
	return m.feeds[id], nil
}

func (m *Model) getFeeds() ([]Feed, error) {
	return m.feeds, nil
}

func sanitize(desc string) string {
	i := strings.Index(desc, "<img")
	if i >= 0 {
		desc = desc[0:i]
	}
	desc = strings.Trim(desc, " \n")
	return desc
}

func (m *Model) refreshFeeds() error {
	rssList, _ := m.readConf()
	id := 0

	for _, src := range rssList {
		response, err := http.Get(src)
		if err != nil {
			return nil
		}

		xdat, err := ioutil.ReadAll(response.Body)
		var rss RSS
		xml.Unmarshal(xdat, &rss)

		for _, item := range rss.Channel.Items {
			item.Desc = sanitize(item.Desc)
			m.feeds = append(m.feeds, Feed{id, rss.Channel.Title, item.Title, item.Desc, item.Link})
			id++
		}
	}

	return nil
}

func (m *Model) run() error {
	m.refreshFeeds()
	return nil
}
