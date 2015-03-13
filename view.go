package main

import (
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"os/exec"
	"strings"
)

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

func displayRSS(feeds []Channel) error {
	w, h, err := terminal.GetSize(int(os.Stdout.Fd()))
	var option int
	newsFeed := []string{""}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	clearScreen()

	numFeeds := len(feeds)
	maxRows := h / 2 //Spaced out by a newline
	wShift := w / 5
	maxPerFeed := (maxRows / numFeeds)
	feedCount := 1
	hdrFmt := fmt.Sprintf("%s%d%s", "%", w/2-3, "s")
	titleFmt := fmt.Sprintf("%s%d%s%d%s", "%", wShift, "d %-", 90, "s %s")

	fmt.Printf(hdrFmt, color.RedString("[MORSEL]")+"\n\n")

	for _, feed := range feeds {
		for i, item := range feed.Items {
			title := strings.Trim(item.Title, " \n")
			title = fmt.Sprintf(titleFmt, feedCount, color.BlueString(title), "["+feed.Title+"]")
			fmt.Printf("%s\n\n", title)
			feedCount++
			newsFeed = append(newsFeed, item.Link)
			if i > maxPerFeed || feedCount >= maxRows-1 {
				break
			}
		}
		if feedCount > maxRows-1 {
			break
		}
	}

	for {
		fmt.Scanln(&option)
		openLink(newsFeed[option])
	}

	return nil
}
