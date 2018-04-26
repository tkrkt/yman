package ui

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/AlecAivazis/survey"
	"github.com/tkrkt/yman/model"
)

func ShowList(manuals []*model.Manual) {
	num := len(manuals)
	if num > 1 {
		fmt.Println("Found " + strconv.Itoa(num) + " manuals")
	} else if num == 0 {
		fmt.Println("No manuals found")
		return
	}

	// show manuals
	head, rows := createList(manuals)
	fmt.Println(head)
	fmt.Println(strings.Join(rows, "\n"))
}

func ShowInteractiveList(manuals []*model.Manual) {
	num := len(manuals)
	if num > 1 {
		fmt.Println("Found " + strconv.Itoa(num) + " manuals")
	} else if num == 0 {
		fmt.Println("No manuals found")
		return
	}

	_, rows := createList(manuals)
	prompt := &survey.Select{
		Message: "Select a manual to show",
		Options: rows,
	}
	var row string
	if err := survey.AskOne(prompt, &row, nil); err != nil {
		Text(err)
		return
	}

	for i, r := range rows {
		if row == r {
			ShowManual(manuals[i], false)
			break
		}
	}
}

func createList(manuals []*model.Manual) (string, []string) {
	buffer := &bytes.Buffer{}
	w := new(tabwriter.Writer)

	w.Init(buffer, 0, 8, 1, ' ', 0)
	fmt.Fprintln(w, "Command\tAuthor\tTitle")
	for _, m := range manuals {
		fmt.Fprintf(w, "%s\t%s\t%s\n", m.Full, m.Author, m.Title)
	}
	w.Flush()

	lines := strings.Split(strings.Trim(buffer.String(), " \n\t"), "\n")
	return lines[0], lines[1:]
}
