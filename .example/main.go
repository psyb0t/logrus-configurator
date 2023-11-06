package main

import (
	_ "github.com/psyb0t/logrus-configurator"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Trace("this shit's a trace")
	logrus.Debug("this shit's a debug")
	logrus.Info("this shit's an info")
	logrus.Warn("this shit's a warn")
	logrus.Error("this shit's an error")
	logrus.Fatal("this shit's a fatal")
}
