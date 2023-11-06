package logrusconfigurator

import (
	"github.com/pkg/errors"
	"github.com/psyb0t/gonfiguration"
	"github.com/sirupsen/logrus"
)

const (
	configKeyLogLevel  = "LOG_LEVEL"
	configKeyLogFormat = "LOG_FORMAT"
	configKeyLogCaller = "LOG_CALLER"
)

const (
	defaultReportCaller = true
	defaultLevel        = levelDebug
	defaultFormat       = formatJSON
)

type config struct {
	Level        level  `env:"LOG_LEVEL"`
	Format       format `env:"LOG_FORMAT"`
	ReportCaller bool   `env:"LOG_CALLER"`
}

func (c config) log() {
	logrus.Debugf(
		"logrus-configurator: level: %s, format: %s, reportCaller: %t",
		c.Level,
		c.Format,
		c.ReportCaller,
	)
}

//nolint:gochecknoinits
func init() {
	if err := configure(); err != nil {
		logrus.Panic(err)
	}
}

func configure() error {
	setDefaults()

	c := config{}
	if err := gonfiguration.Parse(&c); err != nil {
		return errors.Wrap(err, "failed to parse log config")
	}

	if err := setLevel(c.Level); err != nil {
		return errors.Wrap(err, "failed to set log level")
	}

	if err := setFormat(c.Format); err != nil {
		return errors.Wrap(err, "failed to set log format")
	}

	logrus.SetReportCaller(c.ReportCaller)

	c.log()

	return nil
}

func setDefaults() {
	gonfiguration.SetDefaults(map[string]interface{}{
		configKeyLogLevel:  defaultLevel,
		configKeyLogFormat: defaultFormat,
		configKeyLogCaller: defaultReportCaller,
	})
}
