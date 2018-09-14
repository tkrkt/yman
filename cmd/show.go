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
	Short: "Show manuals of the command",
	Long: `Show your manual of the command.
If you want to show your manual of "show" command, use ` + "`yman show show`" + ` instead.`,

	Args: cobra.MaximumNArgs(1),
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

		if query.IsEmpty() {
			cmd.Help()
			return
		}

		// search manuals
		manuals, err := api.Search(query)
		if err != nil {
			ui.Error(err)
			return
		}

		if i, err := cmd.Flags().GetBool("raw"); i && err == nil {
			ui.ShowManuals(manuals, true)
		} else {
			ui.ShowManuals(manuals, false)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().StringP("tag", "t", "", "filter by tag")
	showCmd.Flags().BoolP("raw", "r", false, "show manuals as raw text (without styling)")
}
