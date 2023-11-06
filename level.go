package logrusconfigurator

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type level string

const (
	levelTrace level = "trace"
	levelDebug level = "debug"
	levelInfo  level = "info"
	levelWarn  level = "warn"
	levelError level = "error"
	levelFatal level = "fatal"
	levelPanic level = "panic"
)

func getLogrusLevel(lvl level) (logrus.Level, error) {
	switch lvl {
	case levelTrace:
		return logrus.TraceLevel, nil
	case levelDebug:
		return logrus.DebugLevel, nil
	case levelInfo:
		return logrus.InfoLevel, nil
	case levelWarn:
		return logrus.WarnLevel, nil
	case levelError:
		return logrus.ErrorLevel, nil
	case levelFatal:
		return logrus.FatalLevel, nil
	case levelPanic:
		return logrus.PanicLevel, nil
	default:
		return 0, errors.Wrap(errInvalidLogLevel, string(lvl))
	}
}

func setLevel(lvl level) error {
	logrusLevel, err := getLogrusLevel(lvl)
	if err != nil {
		return err
	}

	logrus.SetLevel(logrusLevel)

	return nil
}
