package ui

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/tkrkt/mdterm"
	"github.com/tkrkt/yman/model"
)

func ShowManual(manual *model.Manual, raw bool) {
	bold := color.New(color.Bold).SprintFunc()

	fmt.Printf("%s (author:%s, tags:%s)\n", bold(manual.Full), manual.Author, manual.Tags)
	if raw {
		fmt.Println(manual.Message)
	} else {
		fmt.Println(string(mdterm.Run([]byte(manual.Message),
			mdterm.WithHeadingStyle(true, 2),
		)))
	}
}

func ShowManuals(manuals []*model.Manual, raw bool) {
	num := len(manuals)
	if num > 1 {
		fmt.Println("Found " + strconv.Itoa(num) + " manuals")
		fmt.Println("")
	} else if num == 0 {
		fmt.Println("No manuals found")
		fmt.Println("")
	}

	// show manuals
	for _, m := range manuals {
		ShowManual(m, raw)
		fmt.Println("")
	}
}
