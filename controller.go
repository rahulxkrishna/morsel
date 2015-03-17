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
	n := c.view.height / 2
	feeds, _ := c.model.getFeeds()

	c.curPos -= n
	if c.curPos < 0 {
		c.curPos = 0
	}
	c.view.displayFeeds(feeds[c.curPos : c.curPos+n])
	c.curPos += n
}

func (c *Controller) getNextFeeds() {
	n := c.view.height / 2

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
	n := c.view.height / 2
	feeds, _ := c.model.getFeeds()
	c.curPos -= 2 * n

	if c.curPos < 0 {
		c.curPos = len(feeds) - n
	}
	c.view.displayFeeds(feeds[c.curPos : c.curPos+n])
	c.curPos += n
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
