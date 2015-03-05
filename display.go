package main

import (
	"fmt"
	tm "github.com/1d4Nf6/goterm"
	"github.com/fatih/color"
)

func displayRSS(items []Item) error {
	tm.Clear()
	box := tm.NewBox(50|tm.PCT, 200, 0)
	for _, item := range items {
		title := color.New(color.FgBlack, color.BgCyan).SprintFunc()
		fmt.Fprint(box, title(item.Title))
		fmt.Fprint(box, item.Desc)
	}
	tm.Print(tm.MoveTo(box.String(), 10|tm.PCT, 10|tm.PCT))
	tm.Flush()
	return nil
}
