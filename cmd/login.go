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
		// check if you are already logged in
		if api.IsLogined() {
			c := api.GetConfig()
			ui.Warn("You are already logged in as " + c.Username)

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

		// login
		email, password, err := ui.Login()
		if err != nil {
			return
		}
		if err = api.Login(email, password); err != nil {
			ui.Error(err)
			return
		}

		c := api.GetConfig()
		ui.Text("Logined as " + c.Username)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
