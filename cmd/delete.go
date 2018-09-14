package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/tkrkt/yman/api"
	"github.com/tkrkt/yman/model"
	"github.com/tkrkt/yman/ui"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [command]",
	Short: "Delete a manual of a command",
	Long: `Delete a manual of a command.
If you want to show your manual of "delete" command, use ` + "`yman show delete`" + ` instead.`,
	Run: func(cmd *cobra.Command, args []string) {
		// create query
		var c string
		if len(args) != 0 {
			c = args[0]
		}
		var tags []string
		if tagString := cmd.Flag("tag").Value.String(); tagString != "" {
			tags = strings.Split(cmd.Flag("tag").Value.String(), ",")
		}
		query := &model.Query{
			Command: c,
			Tags:    tags,
		}

		// search manuals
		manuals, err := api.Search(query)
		if err != nil {
			ui.Error(err)
			return
		}

		m := ui.ShowListForDeletion(manuals)

		if m != nil {
			ui.ShowManual(m, false)
			// delete the manual
			if ans, err := ui.Confirm("Delete this manual?"); ans && err == nil {
				if apiErr := api.Delete(m); apiErr == nil {
					ui.Text("The manual has been deleted")
				} else {
					ui.Error(apiErr)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("tag", "t", "", "select by tag")
}
