package cmd

import (
	"github.com/mds796/CSGY9223-Final/web"
	"github.com/spf13/cobra"
)

var config web.Config
var pidFile string

func init() {
	rootCmd.AddCommand(proxyCmd)

	proxyCmd.PersistentFlags().StringVarP(&pidFile, "pidFile", "p", ".web.pid", "The name of the process ID file.")
	proxyCmd.AddCommand(startCmd)
	proxyCmd.AddCommand(stopCmd)
	proxyCmd.AddCommand(restartCmd)

	setStartArgs(startCmd)
	setStartArgs(restartCmd)
}

func setStartArgs(command *cobra.Command) {
	command.Flags().StringVarP(&config.Host, "host", "H", "localhost", "The host interface to listen on.")
	command.Flags().Uint16VarP(&config.Port, "port", "P", 9898, "The TCP port to listen on.")

	command.Flags().StringVar(&config.UserHost, "destinationHost", "localhost", "The hostname destination file server listens on.")
	command.Flags().Uint16Var(&config.UserPort, "destinationPort", 8989, "The TCP port destination file server listens on.")

	command.Flags().StringVarP(&config.StaticPath, "certificatePath", "C", "./certificate.pem", "The file path to the public certificate for TLS.")
	command.Flags().StringVarP(&config.StaticPath, "keyPath", "K", "./key.pem", "The file path to the secret key for the TLS certificate.")
}

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Runs the proxy server for the secure file server proxy.",
	Long:  `Runs file server proxy that provides security on top of insecure file servers.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the proxy server.",
	Long:  `Starts the proxy server.`,
	Run: func(cmd *cobra.Command, args []string) {
		writePidFile(pidFile)
		web.New(&config).Start()
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the proxy server.",
	Long:  `Stops the proxy server.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopServer(pidFile)
	},
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts the proxy server.",
	Long:  `Restarts the proxy server.`,
	Run: func(cmd *cobra.Command, args []string) {
		stopCmd.Run(cmd, args)
		startCmd.Run(cmd, args)
	},
}
