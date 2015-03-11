package main

import (
	"fmt"
	tm "github.com/1d4Nf6/goterm"
	"github.com/fatih/color"
	"strings"
)

// displayRSS displays the RSS Items in a boxed format
func displayRSS(feeds []Channel) error {
	numFeeds := len(feeds)
	maxPerFeed := MAX_FEEDS / numFeeds
	feedCount := 1

	tm.Clear()
	box := tm.NewBox(80|tm.PCT, 200, 0)

	contentWidth := box.Width - (box.PaddingX+1)*2

	for _, feed := range feeds {
		for i, item := range feed.Items {
			fmtString := fmt.Sprintf("%s%d%s", "%d) %-", contentWidth-contentWidth/4, "s %s")
			title := strings.Trim(fmt.Sprintf(fmtString, feedCount, item.Title, "["+feed.Title+"]"), " \n")
			desc := strings.Trim(fmt.Sprintf("%s", item.Desc), " \n")
			descPrint := color.New(color.FgBlack, color.BgCyan).SprintFunc()
			if len(title) > contentWidth-12 {
				title = title[0 : contentWidth-12]
			}
			if len(desc) > contentWidth-12 {
				desc = desc[0 : contentWidth-12]
			}
			fmt.Fprint(box, title+"\n")
			fmt.Fprint(box, descPrint(desc)+"\n")
			feedCount++

			if i > maxPerFeed {
				break
			}
		}
	}

	tm.Print(tm.MoveTo(box.String(), 10|tm.PCT, 10|tm.PCT))
	tm.Flush()
	return nil
}
