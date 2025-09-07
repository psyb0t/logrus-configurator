package logrusconfigurator

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetDefaults(t *testing.T) {
	unsetEnvs(t)
	setDefaults()
	require.NoError(t, configure(), "Unexpected error")

	actualFormatter := logrus.StandardLogger().Formatter
	defaultFormatter, err := getLogrusFormat(defaultFormat)
	require.NoError(t, err, "Unexpected error")
	assert.IsType(t, defaultFormatter, actualFormatter, "Formatter type mismatch")

	actualLevel := logrus.GetLevel()
	defaultLevel, err := getLogrusLevel(defaultLevel)
	require.NoError(t, err, "Unexpected error")
	assert.Equal(t, defaultLevel, actualLevel, "Log level mismatch")

	actualReportCaller := logrus.StandardLogger().ReportCaller
	assert.Equal(t, defaultReportCaller, actualReportCaller, "ReportCaller mismatch")
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
				require.Error(t, err, "Expected error")

				return
			}

			require.NoError(t, err, "Unexpected error")
			assert.Equal(t, tc.expectedLevel, logrus.GetLevel(), "Log level mismatch")
			assert.IsType(t, tc.expectedFormatter, logrus.StandardLogger().Formatter, "Formatter type mismatch")
		})
	}
}

func TestConfigLog(t *testing.T) {
	testCases := []struct {
		name         string
		level        level
		format       format
		reportCaller bool
	}{
		{
			name:         "Debug config with JSON format and caller reporting",
			level:        levelDebug,
			format:       formatJSON,
			reportCaller: true,
		},
		{
			name:         "Info config with text format and no caller reporting",
			level:        levelInfo,
			format:       formatText,
			reportCaller: false,
		},
		{
			name:         "Error config with JSON format and caller reporting",
			level:        levelError,
			format:       formatJSON,
			reportCaller: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set logrus to debug level to capture debug logs
			originalLevel := logrus.GetLevel()
			defer logrus.SetLevel(originalLevel)
			logrus.SetLevel(logrus.DebugLevel)

			c := config{
				Level:        tc.level,
				Format:       tc.format,
				ReportCaller: tc.reportCaller,
			}

			// This should not panic or error - just exercise the log method
			assert.NotPanics(t, func() {
				c.log()
			}, "config.log() should not panic")
		})
	}
}

func TestConfigureErrorHandling(t *testing.T) {
	testCases := []struct {
		name         string
		logLevel     string
		logFormat    string
		logCaller    string
		expectError  bool
		errorMessage string
	}{
		{
			name:         "Invalid log level",
			logLevel:     "invalid_level",
			logFormat:    "json",
			logCaller:    "false",
			expectError:  true,
			errorMessage: "failed to set log level",
		},
		{
			name:         "Invalid log format",
			logLevel:     "info",
			logFormat:    "invalid_format",
			logCaller:    "false",
			expectError:  true,
			errorMessage: "failed to set log format",
		},
		{
			name:         "Multiple invalid values",
			logLevel:     "invalid_level",
			logFormat:    "invalid_format",
			logCaller:    "false",
			expectError:  true,
			errorMessage: "failed to set log level",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set environment variables
			t.Setenv(configKeyLogLevel, tc.logLevel)
			t.Setenv(configKeyLogFormat, tc.logFormat)
			t.Setenv(configKeyLogCaller, tc.logCaller)

			err := configure()

			if tc.expectError {
				require.Error(t, err, "Expected error for test case: %s", tc.name)
				assert.Contains(t, err.Error(), tc.errorMessage, "Error message should contain expected text")
			} else {
				require.NoError(t, err, "Unexpected error for test case: %s", tc.name)
			}
		})
	}
}
