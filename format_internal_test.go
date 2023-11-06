package logrusconfigurator

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestGetLogrusFormat(t *testing.T) {
	testCases := []struct {
		input       format
		expected    logrus.Formatter
		expectError bool
	}{
		{formatJSON, &logrus.JSONFormatter{}, false},
		{formatText, &logrus.TextFormatter{}, false},
		{"invalid", nil, true},
	}

	for _, tc := range testCases {
		t.Run(string(tc.input), func(t *testing.T) {
			result, err := getLogrusFormat(tc.input)
			if tc.expectError {
				require.Error(t, err)

				return
			}
			require.NoError(t, err)
			require.IsType(t, tc.expected, result)
		})
	}
}

func TestSetFormat(t *testing.T) {
	testCases := []struct {
		format         format
		expectedFormat logrus.Formatter
		expectError    bool
	}{
		{formatJSON, &logrus.JSONFormatter{}, false},
		{formatText, &logrus.TextFormatter{}, false},
		{"Invalid", nil, true},
	}

	for _, tc := range testCases {
		err := setFormat(tc.format)
		if tc.expectError {
			require.NotNil(t, err, "Expected error for format: "+string(tc.format))
		} else {
			require.Nil(t, err, "Unexpected error for format: "+string(tc.format))
			actualFormatter := logrus.StandardLogger().Formatter
			require.IsType(t, tc.expectedFormat, actualFormatter, "Formatter type mismatch")
		}
	}
}
