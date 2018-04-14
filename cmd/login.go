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
		username, password, uierr := ui.Login()
		if uierr != nil {
			return
		}

		// login
		user, apierr := api.Login(username, password)
		if apierr != nil {
			ui.Error(apierr)
			return
		}
		ui.Text("logined as " + user.Username)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
