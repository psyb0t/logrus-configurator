package logrusconfigurator

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func unsetEnvs(t *testing.T) {
	t.Helper()

	require.Nil(t, os.Unsetenv(configKeyLogLevel), "Unexpected error")
	require.Nil(t, os.Unsetenv(configKeyLogFormat), "Unexpected error")
}
