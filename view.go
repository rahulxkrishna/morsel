package main

import (
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/exec"
	"strings"
)

type View struct {
	model *Model
	ctrlr *Controller
}

func (v *View) Init(m *Model, c *Controller) {
	v.model = m
	v.ctrlr = c
}

func (v *View) clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func (v *View) run() {
	var option string
	v.ctrlr.getNextFeeds()
	for {
		fmt.Scanln(&option)
		v.ctrlr.handleInput(option)
		option = ""
	}
}

func (v *View) displayFeeds(feeds []Feed) error {
	w, _, err := terminal.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	v.clearScreen()

	wShift := w / 5
	hdrFmt := fmt.Sprintf("%s%d%s", "%", w/2-3, "s")
	titleFmt := fmt.Sprintf("%s%d%s%d%s", "%", wShift, "d %-", 90, "s %s")

	fmt.Printf(hdrFmt, color.RedString("[MORSEL]")+"\n\n")

	for _, feed := range feeds {
		title := strings.Trim(feed.Title, " \n")
		line := fmt.Sprintf(titleFmt, feed.Id, color.BlueString(title), "["+feed.Source+"]")
		fmt.Printf("%s\n\n", line)
	}
	return nil
}
