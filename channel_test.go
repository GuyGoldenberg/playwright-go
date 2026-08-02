package playwright

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrepareProtocolParams(t *testing.T) {
	credentials := map[string]any{
		"username": "user",
		"password": "secret",
	}
	params, timeout := prepareProtocolParams(map[string]any{
		"timeout":         Float(1234),
		"httpCredentials": credentials,
	})

	require.Equal(t, float64(1234), timeout)
	require.NotContains(t, params, "timeout")
	require.Equal(t, []any{credentials}, params["httpCredentials"])
}

func TestPrepareProtocolParamsPreservesCredentialList(t *testing.T) {
	credentials := []any{map[string]any{"username": "user", "password": "secret"}}
	params, timeout := prepareProtocolParams(map[string]any{
		"httpCredentials": credentials,
	})

	require.Zero(t, timeout)
	require.Equal(t, credentials, params["httpCredentials"])
}
