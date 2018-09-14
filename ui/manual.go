package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mgutz/ansi"
	"github.com/tkrkt/mdterm"
	"github.com/tkrkt/yman/model"
)

// ShowManual displays a manual as styled (if raw is false),
// or as plain text (if raw is true)
func ShowManual(manual *model.Manual, raw bool) {
	// print header
	fmt.Printf("%s%s%s %s%s\n",
		ansi.ColorCode("red+bu")+ansi.DefaultFG,
		manual.Full,
		ansi.Reset+ansi.ColorCode("red+u")+ansi.DefaultFG,
		strings.Join(manual.Tags, ","),
		ansi.Reset,
	)

	// print message
	if raw {
		fmt.Println(manual.Message)
	} else {
		md := string(mdterm.Run([]byte(manual.Message)))
		fmt.Println(strings.Trim(md, " \n\t"))
	}
}

// ShowManuals displays manuals
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
