package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tkrkt/yman/api"
	"github.com/tkrkt/yman/ui"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to yman",
	Long:  `Login to yman.`,

	Run: func(cmd *cobra.Command, args []string) {
		// check if i am already logined
		account, err := api.CurrentAccount()
		if err == nil {
			ui.Text("Already logined as " + account.Username)

			// confirm to logout
			ans, err := ui.Confirm("Login as a different user?")
			if err != nil {
				ui.Error(err)
				return
			}
			if !ans {
				return
			}

			// logout and continue login process
			if err := api.Logout(); err != nil {
				ui.Error(err)
				return
			}
		}

		// get username and password
		username, password, err := ui.Login()
		if err != nil {
			return
		}

		// login
		account, err = api.Login(username, password)
		if err != nil {
			ui.Error(err)
			return
		}
		ui.Text("Logined as " + account.Username)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
