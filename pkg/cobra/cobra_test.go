package cobra

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestInitCmd(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	logrus.SetOutput(&buf)
	defer logrus.SetOutput(logrus.StandardLogger().Out)

	// Run the command
	err := InitCmd()

	// Check for errors
	if err != nil {
		t.Errorf("InitCmd() returned error: %v", err)
	}

	// Check if log output contains expected message
	output := buf.String()
	if !bytes.Contains([]byte(output), []byte("Hello from the base CLI app!")) {
		t.Errorf("Expected log message not found in output: %s", output)
	}
}
