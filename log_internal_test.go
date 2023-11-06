package logrusconfigurator

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestSetDefaults(t *testing.T) {
	unsetEnvs(t)
	setDefaults()
	require.Nil(t, configure(), "Unexpected error")

	actualFormatter := logrus.StandardLogger().Formatter
	defaultFormatter, err := getLogrusFormat(defaultFormat)
	require.Nil(t, err, "Unexpected error")
	require.IsType(t, defaultFormatter, actualFormatter, "Formatter type mismatch")

	actualLevel := logrus.GetLevel()
	defaultLevel, err := getLogrusLevel(defaultLevel)
	require.Nil(t, err, "Unexpected error")
	require.Equal(t, defaultLevel, actualLevel, "Log level mismatch")

	actualReportCaller := logrus.StandardLogger().ReportCaller
	require.Equal(t, defaultReportCaller, actualReportCaller, "ReportCaller mismatch")
}

func TestConfigure(t *testing.T) {
	testCases := []struct {
		name              string
		logLevel          string
		logFormat         string
		expectedLevel     logrus.Level
		expectedFormatter logrus.Formatter
		expectError       bool
	}{
		{
			name:              "Valid level and format",
			logLevel:          "info",
			logFormat:         "json",
			expectedLevel:     logrus.InfoLevel,
			expectedFormatter: &logrus.JSONFormatter{},
			expectError:       false,
		},
		{
			name:        "Invalid level",
			logLevel:    "invalid",
			logFormat:   "json",
			expectError: true,
		},
		{
			name:        "Invalid format",
			logLevel:    "info",
			logFormat:   "invalid",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Setenv(configKeyLogLevel, tc.logLevel)
			t.Setenv(configKeyLogFormat, tc.logFormat)

			err := configure()

			if tc.expectError {
				require.NotNil(t, err, "Expected error")

				return
			}

			require.Nil(t, err, "Unexpected error")
			require.Equal(t, tc.expectedLevel, logrus.GetLevel(), "Log level mismatch")
			require.IsType(t, tc.expectedFormatter, logrus.StandardLogger().Formatter, "Formatter type mismatch")
		})
	}
}
