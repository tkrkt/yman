package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/tkrkt/yman/api"
	"github.com/tkrkt/yman/model"
	"github.com/tkrkt/yman/ui"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show [command]",
	Short: "show manuals of the command",
	Long: `Show your manual of the command.
If you want to show your manual of "show" command, use ` + "`yman show show`" + ` instead.`,

	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// check account (login not required)
		account, err := api.CurrentAccount()
		if err != nil {
			ui.Warn("You are not logged in. Search as a guest user.")
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

		if query.IsEmpty() {
			cmd.Help()
			return
		}

		// search manuals
		manuals, err := api.Search(account, query)
		if err != nil {
			ui.Error(err)
			return
		}

		ui.ShowManuals(manuals)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().StringP("user", "u", "", "filter by username")
	showCmd.Flags().StringP("tag", "t", "", "filter by tag")
}
