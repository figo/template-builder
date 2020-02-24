package cmd

import (
	"github.com/figo/template-builder/pkg/config"
	"github.com/figo/template-builder/pkg/templating"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var templateVars = new(config.TemplateVariables)
var srcRoot string
var dstRoot string

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "templating files",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		if dstRoot == "" {
			path, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			dstRoot = path + "/output"
			os.MkdirAll(dstRoot, 0777)
		}

		templateVars.CompNameCapital = strings.Title(templateVars.CompName)
		templating.WalkDir(srcRoot, dstRoot, templateVars)
		return nil
	},
}

func init() {
	templateCmd.PersistentFlags().StringVar(&templateVars.CompName, "c", "", "component name")
	templateCmd.PersistentFlags().StringVar(&templateVars.MirrorURL, "m", "", "mirror repo URL")
	templateCmd.PersistentFlags().StringVar(&templateVars.UpstreamURL, "u", "", "upstream repo URL")
	templateCmd.PersistentFlags().StringVar(&templateVars.GitRepoOrg, "g", "core-build", "git repository org")
	templateCmd.PersistentFlags().StringVar(&templateVars.BranchName, "b", "branch-0.0.0", "branch name")
	templateCmd.PersistentFlags().StringVar(&templateVars.CompVersion, "o", "0.0.0", "component version")
	templateCmd.PersistentFlags().StringVar(&srcRoot, "s", "", "template root path")
	templateCmd.PersistentFlags().StringVar(&dstRoot, "d", "", "destination root path")
	templateCmd.MarkPersistentFlagRequired("c")
	templateCmd.MarkPersistentFlagRequired("m")
	templateCmd.MarkPersistentFlagRequired("u")
	templateCmd.MarkPersistentFlagRequired("s")
	rootCmd.AddCommand(templateCmd)
}
