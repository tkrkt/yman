package ui

import (
	"fmt"

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
