package proxyutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseConfig(t *testing.T) {
	t.Run("config file", func (t *testing.T) {
		configPath := "../headers.yaml"
		header, err := ParseConfig(configPath)
		require.Nil(t, err)
		fmt.Println(header.UserAgent)
	})
}