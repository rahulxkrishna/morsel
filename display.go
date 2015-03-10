package main

import (
	"fmt"
	tm "github.com/1d4Nf6/goterm"
	"github.com/fatih/color"
	"strings"
)

func displayRSS(items []Item) error {
	tm.Clear()
	box := tm.NewBox(80|tm.PCT, 200, 0)

	contentWidth := box.Width - (box.PaddingX+1)*2

	for _, item := range items {
		title := item.Title
		desc := strings.Trim(item.Desc, "  \n")
		descPrint := color.New(color.FgBlack, color.BgCyan).SprintFunc()
		fmt.Fprint(box, (title))
		if len(desc) > contentWidth-12 {
			desc = desc[0 : contentWidth-12]
			fmt.Fprint(box, descPrint(desc))
		} else {
			fmt.Fprint(box, descPrint(desc))
		}
		fmt.Fprint(box, "\n")
	}
	tm.Print(tm.MoveTo(box.String(), 10|tm.PCT, 10|tm.PCT))
	tm.Flush()
	return nil
}
