package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yusaanthya/cli-base-cobra-logrus/pkg/cobra"
	"github.com/yusaanthya/cli-base-cobra-logrus/pkg/config"
	"github.com/yusaanthya/cli-base-cobra-logrus/pkg/logger"
)

func main() {

	if err := logger.InitLogger(config.LogLevelInfo); err != nil {
		logrus.Fatalf(("Error initializing Logger : %v"), err)
	}

	if err := cobra.InitCmd(); err != nil {
		logrus.Fatalf("Error executing command: %v", err)
		os.Exit(1)
	}

}
