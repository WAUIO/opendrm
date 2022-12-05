package cmd

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	Version string = "unknown"
	Build   string = "unknown"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Elacity DRM",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("------------------------------")
		fmt.Println(cmd.Short)
		fmt.Printf("Version %s\n", Version)
		fmt.Println("------------------------------")
		cmd.Usage()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
