package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A base CLI app with Cobra and logrus",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Hello from the base CLI app!")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("Error executing command: %v", err)
		os.Exit(1)
	}
}
