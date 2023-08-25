package cmd

import (
	"XBS/src"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	server string
	port   string
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server",
	Run: func(cmd *cobra.Command, args []string) {
		os.RemoveAll("middlefile")
		os.MkdirAll("middlefile", os.ModePerm)
		src.Run(server, port)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serverCmd.Flags().StringVarP(&server, "server", "s", "127.0.0.1", "listen addr")
	serverCmd.Flags().StringVarP(&port, "port", "p", "6238", "listen port")
	serverCmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			os.RemoveAll("middlefile")
			os.Exit(1)
		}()
		return nil
	}
}
