package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tkrkt/yman/api"
	"github.com/tkrkt/yman/model"
	"github.com/tkrkt/yman/ui"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [command]",
	Short: "Add your manual of this command",
	Long: `Add your manual for this command.
If you want to show your manual of "add" command, use ` + "`yman show add`" + ` instead.`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// check login status
		account, err := api.CurrentAccount()
		if err != nil {
			ui.Error("You are not logged in. Please login with `yman login` in advance.")
			return
		}

		message := cmd.Flag("message").Value.String()
		if message == "" {
			msg, err := ui.Editor()
			if err != nil {
				ui.Error(err)
				return
			}

			message = msg
		}

		message = strings.Trim(message, " \n\t")

		if message == "" {
			ui.Text("Aborting `yarn add` due to empty message")
			return
		}

		manual := &model.Manual{
			Command:     strings.SplitN(args[0], "/", 2)[0],
			FullCommand: args[0],
			Author:      account.Username,
			Title:       strings.SplitN(message, "\n", 2)[0],
			Message:     message,
			Tags:        strings.Split(cmd.Flag("tag").Value.String(), ","),
		}

		fmt.Println(manual)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("message", "m", "", "content of your manual for this command")
	addCmd.Flags().StringP("tag", "t", "", "tags you want to add to this manual (e.g. -t tag1,tag2)")
}
