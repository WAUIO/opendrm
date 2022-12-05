package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/wauio/opendrm/src/application"
)

var serverPort uint

func init() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.PersistentFlags().UintVarP(&serverPort, "port", "p", 8090, "Server Port")
}

var serveCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Run the license server",
	Run: func(cmd *cobra.Command, args []string) {
		log.WithField("version", Version).Info("Initializing Elacity DRM License Server...")
		application.Bootstrap(fmt.Sprintf("0.0.0.0:%d", serverPort))
	},
}
