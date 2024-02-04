package main

import (
	"github.com/spf13/cobra"
	"golib/cmd/carbon"
)

func init() {
	rootCmd.AddCommand(carbon.Cmd)
}

var rootCmd = &cobra.Command{
	Use:   "???",
	Short: "",
	Run:   nil,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
