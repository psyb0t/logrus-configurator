package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/psyb0t/logrus-configurator"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

func main() {
	logrus.Info("starting with basic config")
	
	// Replace ALL hooks with custom ones
	var customBuffer bytes.Buffer
	customHook := &writer.Hook{
		Writer: io.MultiWriter(os.Stdout, &customBuffer),
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.WarnLevel,
			logrus.ErrorLevel,
		},
	}
	
	logrusconfigurator.SetHooks(customHook)
	
	logrus.Info("this goes to custom hook only")
	logrus.Warn("this shit's a warning through custom hook")
	logrus.Error("this shit's an error through custom hook")
	logrus.Debug("this debug won't show (not in custom hook levels)")
	
	fmt.Printf("\n--- Custom buffer captured: ---\n%s", customBuffer.String())
	fmt.Println("--- End of custom capture ---")
}