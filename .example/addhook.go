package main

import (
	"bytes"
	"fmt"

	"github.com/psyb0t/logrus-configurator"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

func main() {
	logrus.Info("starting with basic config")
	
	// Add a custom hook to capture errors to a buffer
	var errorBuffer bytes.Buffer
	errorHook := &writer.Hook{
		Writer: &errorBuffer,
		LogLevels: []logrus.Level{
			logrus.ErrorLevel,
			logrus.FatalLevel,
			logrus.PanicLevel,
		},
	}
	logrusconfigurator.AddHook(errorHook)
	
	logrus.Info("added custom error hook")
	logrus.Warn("this warning goes to console only")
	logrus.Error("this error goes to both console and buffer")
	
	fmt.Printf("\n--- Error buffer captured: ---\n%s", errorBuffer.String())
	fmt.Println("--- End of captured errors ---")
}