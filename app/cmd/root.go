package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"MyWallet/api/dblayer"
	"MyWallet/api/models"
)

var rootCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate will sync your schema with the database",
	Run: func(cmd *cobra.Command, args []string) {
	  // Do Stuff Here
	  dblayer.DBORM.AutoMigrate(models.FixedCost{})
	},
  }
  
  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Fprintln(os.Stderr, err)
	  os.Exit(1)
	}
  }