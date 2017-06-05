package client

import (
	"crypto/tls"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestResolveTLSConfig(t *testing.T) {
	var timeout = time.Duration(15 * time.Second)
	tr := &http.Transport{TLSHandshakeTimeout: 10 * time.Second,
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{ServerName: "localhost"},
	}
	//defer tr.CloseIdleConnections()
	client := http.Client{
		Timeout:   timeout,
		Transport: tr,
	}
	tlsConfig := resolveTLSConfig(client.Transport)

	require.NotNil(t, tlsConfig)
	require.Equal(t, "localhost", tlsConfig.ServerName)
}