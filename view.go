package main

import (
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/exec"
	"strings"
)

type Feed struct {
	Id     int
	Source string
	Title  string
	Link   string
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func openLink(link string) {
	c := exec.Command("open", link)
	c.Stdout = os.Stdout
	c.Run()
}

func runView() {
	var option string
	getNextFeeds()
	for {
		fmt.Scanln(&option)
		handleInput(option)
	}
}

func displayFeeds(feeds []Feed) error {
	w, _, err := terminal.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	clearScreen()

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
