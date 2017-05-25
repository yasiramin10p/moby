package registry

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestEndpointParse(t *testing.T) {
	testData := []struct {
		str      string
		expected string
	}{
		{IndexServer, IndexServer},
		{"http://0.0.0.0:5000/v1/", "http://0.0.0.0:5000/v1/"},
		{"http://0.0.0.0:5000", "http://0.0.0.0:5000/v1/"},
		{"0.0.0.0:5000", "https://0.0.0.0:5000/v1/"},
		{"http://0.0.0.0:5000/nonversion/", "http://0.0.0.0:5000/nonversion/v1/"},
		{"http://0.0.0.0:5000/v0/", "http://0.0.0.0:5000/v0/v1/"},
	}
	for _, td := range testData {
		e, err := newV1EndpointFromStr(td.str, nil, "", nil)
		if err != nil {
			t.Errorf("%q: %s", td.str, err)
		}
		if e == nil {
			t.Logf("something's fishy, endpoint for %q is nil", td.str)
			continue
		}
		if e.String() != td.expected {
			t.Errorf("expected %q, got %q", td.expected, e.String())
		}
	}
}

func TestEndpointParseInvalid(t *testing.T) {
	testData := []string{
		"http://0.0.0.0:5000/v2/",
	}
	for _, td := range testData {
		e, err := newV1EndpointFromStr(td, nil, "", nil)
		if err == nil {
			t.Errorf("expected error parsing %q: parsed as %q", td, e)
		}
	}
}

func TestNewV1Endpoint(t *testing.T)  {
	uri, error := url.Parse("http://0.0.0.0:5000/v2/")
	assert.NoError(t, error)
	var index = makeIndex("/v1/")
	tlsConfig, error := newTLSConfig(index.Name, index.Secure)
	assert.NoError(t, error)
	v1EndPoint, error := newV1Endpoint(*uri,tlsConfig,"",nil)
	require.NotNil(t,v1EndPoint)
	assert.NoError(t, error)
	require.EqualValues(t,v1EndPoint.URL.Host,uri.Host)
	require.EqualValues(t,v1EndPoint.IsSecure,index.Secure)
}
// Ensure that a registry endpoint that responds with a 401 only is determined
// to be a valid v1 registry endpoint
func TestValidateEndpoint(t *testing.T) {
	requireBasicAuthHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("WWW-Authenticate", `Basic realm="localhost"`)
		w.WriteHeader(http.StatusUnauthorized)
	})

	// Make a test server which should validate as a v1 server.
	testServer := httptest.NewServer(requireBasicAuthHandler)
	defer testServer.Close()

	testServerURL, err := url.Parse(testServer.URL)
	if err != nil {
		t.Fatal(err)
	}

	testEndpoint := V1Endpoint{
		URL:    testServerURL,
		client: HTTPClient(NewTransport(nil)),
	}

	if err = validateEndpoint(&testEndpoint); err != nil {
		t.Fatal(err)
	}

	if testEndpoint.URL.Scheme != "http" {
		t.Fatalf("expecting to validate endpoint as http, got url %s", testEndpoint.String())
	}
}
