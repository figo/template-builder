package cmd

import (
	"github.com/spf13/cobra"
)

var compName string
var mirrorURL string
var upstreamURL string
var templateDir string

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "templating files",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	templateCmd.PersistentFlags().StringVar(&compName, "compName", "", "component name")
	templateCmd.PersistentFlags().StringVar(&mirrorURL, "mirrorURL", "", "mirror repo URL")
	templateCmd.PersistentFlags().StringVar(&upstreamURL, "upstreamURL", "", "upstream repo URL")
	templateCmd.MarkPersistentFlagRequired("compName")
	templateCmd.MarkPersistentFlagRequired("mirrorURL")
	templateCmd.MarkPersistentFlagRequired("upstreamURL")
	rootCmd.AddCommand(templateCmd)
}
