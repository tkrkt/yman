package cmd

import (
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tkrkt/yman/api"
	"github.com/tkrkt/yman/model"
	"github.com/tkrkt/yman/ui"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [command]",
	Short: "Add your manual of a command",
	Long: `Add your manual of a command.
If you want to show your manual of "add" command, use ` + "`yman show add`" + ` instead.`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// check login status (login required)
		if !api.IsLogined() {
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

		// extract title as first line of message with removing # (headings)
		reg, _ := regexp.Compile("^#*\\s*")
		title := strings.SplitN(message, "\n", 2)[0]
		title = reg.ReplaceAllString(title, "")

		var tags []string
		if tagString := cmd.Flag("tag").Value.String(); tagString != "" {
			tags = strings.Split(cmd.Flag("tag").Value.String(), ",")
		}

		config := api.GetConfig()
		manual := &model.Manual{
			Command: strings.SplitN(args[0], "/", 2)[0],
			Full:    args[0],
			Author:  config.Username,
			Title:   title,
			Message: message,
			Tags:    tags,
		}

		api.Add(manual)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("message", "m", "", "content of your manual for this command")
	addCmd.Flags().StringP("tag", "t", "", "tags you want to add to this manual (e.g. -t tag1,tag2)")
}
