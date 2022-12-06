package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number and build",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("elacity-drm version %s, build %s", Version, Build))
		fmt.Println(fmt.Sprintf("go version %s", runtime.Version()))
	},
}
