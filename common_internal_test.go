package logrusconfigurator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func unsetEnvs(t *testing.T) {
	t.Helper()

	require.NoError(t, os.Unsetenv(configKeyLogLevel), "Unexpected error")
	require.NoError(t, os.Unsetenv(configKeyLogFormat), "Unexpected error")
}
