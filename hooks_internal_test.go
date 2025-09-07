package logrusconfigurator

import (
	"bytes"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetLoggerHooks(t *testing.T) {
	logger := logrus.New()
	buffer1 := &bytes.Buffer{}
	buffer2 := &bytes.Buffer{}

	hook1 := getStderrHook(buffer1)
	hook2 := getStdoutHook(buffer2)

	setLoggerHooks(logger, hook1, hook2)

	assert.Len(t, logger.Hooks, 7, "Expected 7 hook levels")
	
	// Verify hooks are properly set for stderr levels
	stderrLevels := []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
	}
	for _, level := range stderrLevels {
		assert.Len(t, logger.Hooks[level], 1, "Expected 1 hook for level %v", level)
	}
	
	// Verify hooks are properly set for stdout levels
	stdoutLevels := []logrus.Level{
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	}
	for _, level := range stdoutLevels {
		assert.Len(t, logger.Hooks[level], 1, "Expected 1 hook for level %v", level)
	}
}

func TestSetHooks(t *testing.T) {
	originalHooks := logrus.StandardLogger().Hooks
	defer func() {
		logrus.StandardLogger().Hooks = originalHooks
	}()

	buffer1 := &bytes.Buffer{}
	buffer2 := &bytes.Buffer{}
	
	hook1 := getStderrHook(buffer1)
	hook2 := getStdoutHook(buffer2)

	SetHooks(hook1, hook2)

	assert.Len(t, logrus.StandardLogger().Hooks, 7, "Expected 7 hook levels")
}

func TestAddHook(t *testing.T) {
	originalLogger := logrus.StandardLogger()
	defer func() {
		logrus.SetOutput(originalLogger.Out)
		logrus.StandardLogger().Hooks = originalLogger.Hooks
	}()

	// Clear existing hooks
	clearLoggerHooks(logrus.StandardLogger())

	buffer := &bytes.Buffer{}
	hook := getStderrHook(buffer)

	AddHook(hook)

	stderrLevels := []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
	}
	
	for _, level := range stderrLevels {
		assert.Len(t, logrus.StandardLogger().Hooks[level], 1, "Expected 1 hook for level %v", level)
	}
}

func TestGetStderrHookWithCustomWriter(t *testing.T) {
	buffer := &bytes.Buffer{}
	hook := getStderrHook(buffer)

	require.NotNil(t, hook, "Hook should not be nil")
	
	writerHook, ok := hook.(*writer.Hook)
	require.True(t, ok, "Hook should be of type writer.Hook")
	
	assert.Equal(t, buffer, writerHook.Writer, "Writer should match provided buffer")
	
	expectedLevels := []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
	}
	assert.Equal(t, expectedLevels, writerHook.LogLevels, "Log levels should match expected stderr levels")
}

func TestGetStdoutHookWithCustomWriter(t *testing.T) {
	buffer := &bytes.Buffer{}
	hook := getStdoutHook(buffer)

	require.NotNil(t, hook, "Hook should not be nil")
	
	writerHook, ok := hook.(*writer.Hook)
	require.True(t, ok, "Hook should be of type writer.Hook")
	
	assert.Equal(t, buffer, writerHook.Writer, "Writer should match provided buffer")
	
	expectedLevels := []logrus.Level{
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	}
	assert.Equal(t, expectedLevels, writerHook.LogLevels, "Log levels should match expected stdout levels")
}