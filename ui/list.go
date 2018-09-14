package ui

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/AlecAivazis/survey"
	"github.com/mgutz/ansi"
	"github.com/tkrkt/yman/model"
)

// ShowList shows list of manuals
func ShowList(manuals []*model.Manual) {
	num := len(manuals)
	if num > 1 {
		fmt.Println("Found " + strconv.Itoa(num) + " manuals")
	} else if num == 0 {
		fmt.Println(ansi.Red + "No manuals found" + ansi.Reset)
		return
	}

	// show manuals
	head, rows := createList(manuals)
	fmt.Println(ansi.ColorCode("cyan") + head + ansi.Reset)
	fmt.Println(strings.Join(rows, "\n"))
}

// ShowInteractiveList displays list of manuals
// and displayss cursor for selecting a manual, and displays it if selected.
func ShowInteractiveList(manuals []*model.Manual) {
	num := len(manuals)
	if num > 1 {
		fmt.Println("Found " + strconv.Itoa(num) + " manuals")
	} else if num == 0 {
		fmt.Println(ansi.Red + "No manuals found" + ansi.Reset)
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

// ShowListForDeletion displays list of manuals and displays cursor for deletion.
func ShowListForDeletion(manuals []*model.Manual) *model.Manual {
	num := len(manuals)
	if num > 1 {
		fmt.Println("Found " + strconv.Itoa(num) + " manuals")
	} else if num == 0 {
		fmt.Println(ansi.Red + "No manuals found" + ansi.Reset)
		return nil
	}

	var manual *model.Manual

	if num > 1 {
		_, rows := createList(manuals)
		prompt := &survey.Select{
			Message: "Select the manual to delete",
			Options: rows,
		}
		var row string
		if err := survey.AskOne(prompt, &row, nil); err != nil {
			Text(err)
			return nil
		}

		for i, r := range rows {
			if row == r {
				manual = manuals[i]
				break
			}
		}
	} else { // num === 1
		manual = manuals[0]
	}

	if manual == nil {
		return nil
	}
	return manual
}

func createList(manuals []*model.Manual) (string, []string) {
	buffer := &bytes.Buffer{}
	w := new(tabwriter.Writer)

	w.Init(buffer, 0, 8, 1, ' ', 0)
	fmt.Fprintln(w, "Command\tTitle")
	for _, m := range manuals {
		fmt.Fprintf(w, "%s\t%s\n", m.Full, m.Title)
	}
	w.Flush()

	lines := strings.Split(strings.Trim(buffer.String(), " \n\t"), "\n")
	return lines[0], lines[1:]
}
