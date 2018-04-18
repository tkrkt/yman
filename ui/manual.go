package ui

import (
	"fmt"
	"strconv"

	"github.com/tkrkt/yman/model"
)

func ShowManual(manual *model.Manual) {
	fmt.Println("--------")
	fmt.Println(manual.FullCommand)
	fmt.Println("author:", manual.Author)
	if len(manual.Tags) > 0 {
		fmt.Println("tags:", manual.Tags)
	}
	fmt.Println("--")
	fmt.Println(manual.Message)
}

func ShowManuals(manuals []*model.Manual) {
	num := len(manuals)
	if num > 1 {
		fmt.Println("Found " + strconv.Itoa(num) + " manuals")
	} else if num == 0 {
		fmt.Println("No manuals found")
	}

	// show manuals
	for _, m := range manuals {
		ShowManual(m)
	}
}
