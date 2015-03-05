package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	_ "github.com/jroimartin/gocui"
	"io/ioutil"
	_ "log"
	"net/http"
	"os"
	_ "time"
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

func readConf() ([]string, error) {
	file, err := os.Open("feeder.conf")
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

	for _, src := range rssList {
		response, err := http.Get(src)
		if err != nil {
			return rssText, err
		}

		xdat, err := ioutil.ReadAll(response.Body)
		var rss RSS
		xml.Unmarshal(xdat, &rss)

		for i, item := range rss.Channel.Items {
			title := fmt.Sprintf("%d) %s\n", i+1, item.Title)
			desc := fmt.Sprintf("\t. %s\n", item.Desc)
			item := Item{title, desc, "", ""}
			rssText = append(rssText, item)
			if i > 10 {
				break
			}
		}
	}

	return rssText, nil
}

func main() {
	for {
		rssList, _ := fetchRSS()
		displayRSS(rssList)
		break
		break
		//time.Sleep(5000 * time.Millisecond)
	}
}
