package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "yman",
	Short: "Manual management tool for CLI applications",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	// insert command "show" in args if no subcommand is found
	_, _, err := rootCmd.Find(os.Args[1:])
	if err != nil && strings.HasPrefix(err.Error(), "unknown command") {
		os.Args = append(os.Args[:1], append([]string{"show"}, os.Args[1:]...)...)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}
