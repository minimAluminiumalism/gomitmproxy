package proxyutil

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestParseConfig(t *testing.T) {
	t.Run("config file", func (t *testing.T) {
		configPath := "../headers.yaml"
		_, err := ParseConfig(configPath)
		require.Nil(t, err)
	})
}