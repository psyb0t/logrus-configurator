package logrusconfigurator

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestGetLogrusLevel(t *testing.T) {
	testCases := []struct {
		input       level
		expected    logrus.Level
		expectError bool
	}{
		{levelTrace, logrus.TraceLevel, false},
		{levelDebug, logrus.DebugLevel, false},
		{levelInfo, logrus.InfoLevel, false},
		{levelWarn, logrus.WarnLevel, false},
		{levelError, logrus.ErrorLevel, false},
		{levelFatal, logrus.FatalLevel, false},
		{levelPanic, logrus.PanicLevel, false},
		{"invalid", 0, true},
	}

	for _, tc := range testCases {
		t.Run(string(tc.input), func(t *testing.T) {
			result, err := getLogrusLevel(tc.input)
			if tc.expectError {
				require.Error(t, err)

				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestSetLevel(t *testing.T) {
	testCases := []struct {
		level         level
		expectedLevel logrus.Level
		expectError   bool
	}{
		{levelTrace, logrus.TraceLevel, false},
		{levelDebug, logrus.DebugLevel, false},
		{levelInfo, logrus.InfoLevel, false},
		{levelWarn, logrus.WarnLevel, false},
		{levelError, logrus.ErrorLevel, false},
		{levelFatal, logrus.FatalLevel, false},
		{levelPanic, logrus.PanicLevel, false},
		{"Invalid", logrus.PanicLevel, true},
	}

	for _, tc := range testCases {
		err := setLevel(tc.level)
		if tc.expectError {
			require.NotNil(t, err, "Expected error for level: "+string(tc.level))
		} else {
			require.Nil(t, err, "Unexpected error for level: "+string(tc.level))
			require.Equal(t, tc.expectedLevel, logrus.GetLevel(), "Log level mismatch")
		}
	}
}
