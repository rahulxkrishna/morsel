package main

import (
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/exec"
	"strings"
)

type MView interface {
	Init(m *Model, c *Controller)
	Run()
	DisplayFeeds(feeds []Feed) error
	Maxlines() int
}

type CLView struct {
	model  *Model
	ctrlr  *Controller
	height int
	width  int
}

const head = "[MORSEL]"

func (v *CLView) getTermSz() {
	w, h, err := terminal.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		fmt.Println("Using default values")
		v.width = 200
		v.height = 50
	} else {
		v.width = w
		v.height = h - 4
	}
}

func (v *CLView) Init(m *Model, c *Controller) {
	v.model = m
	v.ctrlr = c
	v.getTermSz()
}

func (v *CLView) clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func (v *CLView) Run() {
	var option string
	v.ctrlr.getNextFeeds()
	for {
		fmt.Scanln(&option)
		v.ctrlr.handleInput(option)
		option = ""
	}
}

func (v *CLView) Maxlines() int {
	return v.height
}

func (v *CLView) DisplayFeeds(feeds []Feed) error {
	left := 0
	right := 0

	v.clearScreen()

	if v.width > 110 {
		left = v.width / 5
		right = 100
	}

	hdrFmt := fmt.Sprintf("%s%d%s", "%", v.width/2+len(head)/2, "s")
	titleFmt := fmt.Sprintf("%s%d%s%d%s", "%", left, "d %-", right, "s %s")

	fmt.Printf(hdrFmt, color.RedString(head)+"\n\n")

	for _, feed := range feeds {
		title := strings.Trim(feed.Title, " \n")
		line := fmt.Sprintf(titleFmt, feed.Id, color.BlueString(title), "["+feed.Source+"]")
		fmt.Printf("%s\n\n", line)
	}
	return nil
}
