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
		// check login status (login required)
		account, err := api.CurrentAccount()
		if err != nil {
			ui.Error("You are not logged in. Please login with `yman login` in advance.")
			return
		}

		// create query
		var c string
		if len(args) != 0 {
			c = args[0]
		}
		author := cmd.Flag("user").Value.String()
		var tags []string
		if tagString := cmd.Flag("tag").Value.String(); tagString != "" {
			tags = strings.Split(cmd.Flag("tag").Value.String(), ",")
		}
		query := &model.Query{
			Command: c,
			Author:  author,
			Tags:    tags,
		}

		// search manuals
		manuals, err := api.Search(account, query)
		if err != nil {
			ui.Error(err)
			return
		}

		m := ui.ShowListForDeletion(manuals)

		if m != nil {
			ui.ShowManual(m, false)
			if m.Author == account.Username {
				// delete the manual you created
				if ans, err := ui.Confirm("Delete this manual?"); ans && err == nil {
					if apiErr := api.Delete(account, m); apiErr == nil {
						ui.Text("The manual has been deleted")
					} else {
						ui.Error(apiErr)
					}
				}
			} else {
				// unstock the manual created by others
				if ans, err := ui.Confirm("Unstock this manual?"); ans && err == nil {
					api.Unstock(account, m)
					if apiErr := api.Unstock(account, m); apiErr == nil {
						ui.Text("The manual is unstocked")
					} else {
						ui.Error(apiErr)
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().StringP("user", "u", "", "filter by username")
	deleteCmd.Flags().StringP("tag", "t", "", "filter by tag")
}
