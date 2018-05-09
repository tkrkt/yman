package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tkrkt/yman/api"
	"github.com/tkrkt/yman/ui"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from yman",
	Long:  `Logout from yman.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !api.IsLogined() {
			ui.Error("You have already logged out.")
			return
		}

		if err := api.Logout(); err != nil {
			ui.Error(err)
			return
		}

		ui.Text("Logged out successfully.")
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
