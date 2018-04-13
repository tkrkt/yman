package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tkrkt/yman/api"
	"github.com/tkrkt/yman/ui"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check if i am already logined
		currentUser := api.CurrentUser()
		if currentUser != nil {
			ui.Text("already logined as " + currentUser.Username)
			return
		}

		// get username and password
		username, password, err := ui.Login()
		if err != nil {
			return
		}

		// // login
		api.Login(username, password)

		// write access token into ~/.ymanrc
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
