package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/tkrkt/yman/api"
	"github.com/tkrkt/yman/model"
	"github.com/tkrkt/yman/ui"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [command]",
	Short: "list created or stocked manuals",
	Long:  `list created or stocked manuals`,

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

		// search manuals
		manuals, err := api.Search(query)
		if err != nil {
			ui.Error(err)
			return
		}
		if i, err := cmd.Flags().GetBool("interactive"); i && err == nil {
			ui.ShowInteractiveList(manuals)
		} else {
			ui.ShowList(manuals)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("tag", "t", "", "filter by tag")
	listCmd.Flags().BoolP("interactive", "i", false, "enable to select a manual by cursor")
}
