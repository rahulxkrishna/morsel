package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

const MAX_FEEDS = 25

type Controller struct {
	curPos int
	model  *Model
	view   *View
}

func (c *Controller) Init(m *Model, v *View) {
	c.model = m
	c.view = v
}

func (c *Controller) refreshFeeds() {
	feeds, _ := c.model.getFeeds()
	c.view.displayFeeds(feeds[c.curPos-MAX_FEEDS : c.curPos])
}

func (c *Controller) getNextFeeds() {
	n := MAX_FEEDS

	feeds, _ := c.model.getFeeds()

	if c.curPos >= len(feeds) {
		c.curPos = 0
	}

	if c.curPos+n > len(feeds) {
		n = len(feeds) - c.curPos
	}

	c.view.displayFeeds(feeds[c.curPos : c.curPos+n])
	c.curPos += n
}

func (c *Controller) getPrevFeeds() {
	//displayFeeds(morsels)
}

func (c *Controller) openLink(ip string) {
	id, _ := strconv.Atoi(ip[1:])
	feed, _ := c.model.getFeed(id)
	cmd := exec.Command("open", feed.Link)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (c *Controller) detailFeed(ip string) {
	id, _ := strconv.Atoi(ip[1:])
	feed, _ := c.model.getFeed(id)
	fmt.Println(feed.Desc)
}

func (c *Controller) handleInput(ip string) {
	if ip == "" {
		c.getNextFeeds()
		return
	}

	switch ip[0] {
	case 'n':
		c.getNextFeeds()
	case 'p':
		c.getPrevFeeds()
	case 'o':
		c.openLink(ip)
	case 'd':
		c.detailFeed(ip)
	case 'r':
		c.refreshFeeds()
	}
}
