package cmd

import (
	"github.com/spf13/cobra"

	"github.com/wauio/opendrm/src/cmd/tool"
)

func init() {
	toolCmd.AddCommand(tool.WidevineSign)
	toolCmd.AddCommand(tool.WidevineGenerate)
	rootCmd.AddCommand(toolCmd)
}

var toolCmd = &cobra.Command{
	Use:   "tool",
	Short: "Set of tools to manage, sandbox, check, monitoring license management",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}
