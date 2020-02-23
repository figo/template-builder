package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
}

var rootCmd = &cobra.Command{
	Use:   "template-builder",
	Short: "templating files",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		return
	},
}

// Execute run the command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
