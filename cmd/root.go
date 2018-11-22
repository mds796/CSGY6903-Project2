package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	"github.com/spf13/cobra"
)

var profilerFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "CSGY6903-Project2",
	Short: "A secure HTTPS proxy for file servers.",
	Long:  `An HTTPS proxy that provides secure storage on top of an insecure file server.`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initProfiler)

	// See https://blog.golang.org/profiling-go-programs
	rootCmd.PersistentFlags().StringVar(&profilerFile, "profile", "./profile.prof", "profiler file (default is ./profile.prof)")
}

func initProfiler() {
	if profilerFile != "" {
		f, err := os.Create(profilerFile)
		if err != nil {
			log.Fatal(err)
		}

		if err = pprof.StartCPUProfile(f); err != nil {
			log.Printf("Unable to start CPU profile. Error: %v", err)
		}

		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		go func() {
			<-sigs
			pprof.StopCPUProfile()
			os.Exit(0)
		}()
	}
}
