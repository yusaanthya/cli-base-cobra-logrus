package cobra

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/yusaanthya/cli-base-cobra-logrus/pkg/config"
)

func InitCmd() error {
	rootCmd := &cobra.Command{
		Use:   "app",
		Short: "A base CLI app with Cobra and logrus",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.Info("Hello from the base CLI app!")
		},
	}

	// Add version command
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version:    %s\n", config.AppVersion)
			fmt.Printf("Git Commit: %s\n", config.GitCommit)
			fmt.Printf("Build Date: %s\n", config.BuildDate)
			fmt.Printf("Environment: %s\n", config.AppEnv)
		},
	}
	rootCmd.AddCommand(versionCmd)

	err := rootCmd.Execute()
	if err != nil {
		logrus.Fatalf("Error executing command: %v", err)
	}

	return err
}
