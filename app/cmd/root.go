package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "mywallet",
	Short: "MyWallet API server",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("You need to specify if you want to run the server or migrations")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
