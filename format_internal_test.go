package logrusconfigurator

import (
	"runtime"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
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
			require.Error(t, err, "Expected error for format: "+string(tc.format))
		} else {
			require.NoError(t, err, "Unexpected error for format: "+string(tc.format))
			actualFormatter := logrus.StandardLogger().Formatter
			assert.IsType(t, tc.expectedFormat, actualFormatter, "Formatter type mismatch")
		}
	}
}

func TestCallerPrettyfier(t *testing.T) {
	testCases := []struct {
		name   string
		format format
	}{
		{
			name:   "JSON format with caller prettyfier",
			format: formatJSON,
		},
		{
			name:   "Text format with caller prettyfier",
			format: formatText,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			formatter, err := getLogrusFormat(tc.format)
			require.NoError(t, err)

			// Create a mock runtime frame
			frame := &runtime.Frame{
				Function: "github.com/example/package.TestFunction",
				File:     "/path/to/file/test.go",
				Line:     42,
			}

			var funcName, fileName string
			
			switch f := formatter.(type) {
			case *logrus.JSONFormatter:
				require.NotNil(t, f.CallerPrettyfier, "CallerPrettyfier should be set for JSON formatter")
				funcName, fileName = f.CallerPrettyfier(frame)
			case *logrus.TextFormatter:
				require.NotNil(t, f.CallerPrettyfier, "CallerPrettyfier should be set for Text formatter")
				funcName, fileName = f.CallerPrettyfier(frame)
			default:
				t.Fatalf("Unexpected formatter type: %T", formatter)
			}

			// Verify the function name format
			assert.Equal(t, "github.com/example/package.TestFunction()", funcName, "Function name should end with ()")
			
			// Verify the file name format
			assert.True(t, strings.HasSuffix(fileName, "test.go:42"), "File name should contain base filename and line number")
			assert.Equal(t, "test.go:42", fileName, "File name should be formatted as basename:line")
		})
	}
}

func TestCallerPrettyfierFormatsCorrectly(t *testing.T) {
	formatter, err := getLogrusFormat(formatJSON)
	require.NoError(t, err)

	jsonFormatter := formatter.(*logrus.JSONFormatter)
	
	testCases := []struct {
		name         string
		frame        *runtime.Frame
		expectedFunc string
		expectedFile string
	}{
		{
			name: "Standard function call",
			frame: &runtime.Frame{
				Function: "main.main",
				File:     "/home/user/project/main.go",
				Line:     10,
			},
			expectedFunc: "main.main()",
			expectedFile: "main.go:10",
		},
		{
			name: "Package function call",
			frame: &runtime.Frame{
				Function: "github.com/user/repo/pkg.Function",
				File:     "/go/src/github.com/user/repo/pkg/file.go",
				Line:     25,
			},
			expectedFunc: "github.com/user/repo/pkg.Function()",
			expectedFile: "file.go:25",
		},
		{
			name: "Method call",
			frame: &runtime.Frame{
				Function: "(*MyStruct).Method",
				File:     "/path/to/struct.go",
				Line:     100,
			},
			expectedFunc: "(*MyStruct).Method()",
			expectedFile: "struct.go:100",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			funcName, fileName := jsonFormatter.CallerPrettyfier(tc.frame)
			assert.Equal(t, tc.expectedFunc, funcName, "Function name mismatch")
			assert.Equal(t, tc.expectedFile, fileName, "File name mismatch")
		})
	}
}
