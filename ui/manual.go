package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/tkrkt/mdterm"
	"github.com/tkrkt/yman/model"
)

func ShowManual(manual *model.Manual, raw bool) {
	bold := color.New(color.Bold).SprintFunc()
	// print header

	fmt.Printf("%s (author:%s, tags:%s)\n", bold(manual.Full), manual.Author, manual.Tags)
	// print message
	if raw {
		fmt.Println(manual.Message)
	} else {
		md := string(mdterm.Run([]byte(manual.Message)))
		fmt.Println(strings.Trim(md, " \n\t"))
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
