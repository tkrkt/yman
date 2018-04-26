package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mgutz/ansi"
	"github.com/tkrkt/mdterm"
	"github.com/tkrkt/yman/model"
)

func ShowManual(manual *model.Manual, raw bool) {
	// print header
	if len(manual.Tags) > 0 {
		fmt.Printf("%s%s%s [author:%s, tags:%s]%s\n",
			ansi.ColorCode("red+bu")+ansi.DefaultFG,
			manual.Full,
			ansi.Reset+ansi.ColorCode("red+u")+ansi.DefaultFG,
			manual.Author,
			strings.Join(manual.Tags, ","),
			ansi.Reset,
		)
	} else {
		fmt.Printf("%s%s%s [author:%s]%s\n",
			ansi.ColorCode("red+bu")+ansi.DefaultFG,
			manual.Full,
			ansi.Reset+ansi.ColorCode("red+u")+ansi.DefaultFG,
			manual.Author,
			ansi.Reset,
		)
	}

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
