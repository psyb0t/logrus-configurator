package logrusconfigurator

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type format string

const (
	formatJSON format = "json"
	formatText format = "text"
)

func getLogrusFormat(fmt format) (logrus.Formatter, error) { //nolint:ireturn
	switch fmt {
	case formatJSON:
		return &logrus.JSONFormatter{}, nil
	case formatText:
		return &logrus.TextFormatter{}, nil
	default:
		return nil, errors.Wrap(errInvalidLogFormat, string(fmt))
	}
}

func setFormat(fmt format) error {
	logrusFormatter, err := getLogrusFormat(fmt)
	if err != nil {
		return err
	}

	logrus.SetFormatter(logrusFormatter)

	return nil
}
