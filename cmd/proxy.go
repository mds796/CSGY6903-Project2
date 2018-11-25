package cmd

import (
	"github.com/mds796/CSGY6903-Project2/proxy"
	"github.com/spf13/cobra"
	"log"
)

var config proxy.Config
var pidFile string

func init() {
	rootCmd.AddCommand(proxyCmd)

	proxyCmd.PersistentFlags().StringVarP(&pidFile, "pidFile", "p", ".proxy.pid", "The name of the process ID file.")
	proxyCmd.AddCommand(startCmd)
	proxyCmd.AddCommand(stopCmd)
	proxyCmd.AddCommand(restartCmd)

	setStartArgs(startCmd)
	setStartArgs(restartCmd)
}

func setStartArgs(command *cobra.Command) {
	command.Flags().StringVarP(&config.Host, "host", "H", "localhost", "The host interface to listen on.")
	command.Flags().Uint16VarP(&config.Port, "port", "P", 9898, "The TCP port to listen on.")

	command.Flags().StringVar(&config.DestinationScheme, "destinationScheme", "http", "The protocol scheme destination file server listens on.")
	command.Flags().StringVar(&config.DestinationHost, "destinationHost", "localhost", "The hostname destination file server listens on.")
	command.Flags().Uint16Var(&config.DestinationPort, "destinationPort", 8989, "The TCP port destination file server listens on.")

	command.Flags().StringVarP(&config.SymmetricKeyPath, "symmetricKeyPath", "S", "./symmetric.key", "The file path to the symmetric key for file encryption.")
	command.Flags().StringVarP(&config.CertificatePath, "certificatePath", "C", "./server.crt", "The file path to the public certificate for TLS.")
	command.Flags().StringVarP(&config.KeyPath, "keyPath", "K", "./server.key", "The file path to the secret key for the TLS certificate.")

	command.Flags().StringVarP(&config.UploadApi, "upload", "U", "/!/upload", "The URL path to upload a file to the destination server.")
	command.Flags().StringVarP(&config.DownloadApi, "download", "D", "/!/dl/", "The URL path to download a file from the destination server.")
	command.Flags().StringVarP(&config.WebSocketApi, "websocket", "W", "/!/socket", "The URL path to start a Web Socket with the destination server.")
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

		if err := proxy.NewProxy(&config).Start(); err != nil {
			log.Fatalln(err)
		}
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
