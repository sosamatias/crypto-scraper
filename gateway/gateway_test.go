package gateway

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const apiKeyMocked = "apiKeyMocked"

func Test_Gateway_OK(t *testing.T) {
	t.Parallel()

	serverMocked := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		d, err := ioutil.ReadFile("data_test/data_test.json")
		assert.NoError(t, err)
		res.Write(d)
	}))
	defer serverMocked.Close()

	gateway, err := NewGateway(apiKeyMocked, serverMocked.URL, serverMocked.Client())
	assert.NoError(t, err)

	r, err := gateway.List(SortMarketCap)
	assert.Equal(t, 100, len(r.Data))
	assert.NoError(t, err)
}

func Test_Gateway_Error_StatusCode500(t *testing.T) {
	t.Parallel()

	serverMocked := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		req.Response.StatusCode = 500
	}))
	defer serverMocked.Close()

	gateway, err := NewGateway(apiKeyMocked, serverMocked.URL, serverMocked.Client())
	assert.NoError(t, err)

	_, err = gateway.List(SortMarketCap)
	assert.Error(t, err)
}

func Test_Gateway_Error_Body(t *testing.T) {
	t.Parallel()

	serverMocked := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		// This handler tells it has 1 byte body,
		// but actually it sends none.
		// So at the other end (the client) when attempting to read 1 byte from it,
		// obviously that won't succeed, and will result in the following error:
		// Unable to read from body unexpected EOF
		res.Header().Set("Content-Length", "1")
	}))
	defer serverMocked.Close()

	gateway, err := NewGateway(apiKeyMocked, serverMocked.URL, serverMocked.Client())
	assert.NoError(t, err)

	_, err = gateway.List(SortMarketCap)
	assert.Error(t, err)
}

func Test_Gateway_Error_BaseURL(t *testing.T) {
	t.Parallel()

	const baseURL = "$#$%%^"
	gateway, err := NewGateway(apiKeyMocked, baseURL, nil)
	assert.NoError(t, err)

	_, err = gateway.List(SortMarketCap)
	assert.Error(t, err)
}

func Test_NewGateway_ErrAPIKeyRequired(t *testing.T) {
	t.Parallel()

	const apiKeyEmpty = ""
	_, err := NewGateway(apiKeyEmpty, "", nil)
	assert.Equal(t, ErrAPIKeyRequired, err)
}

func Test_NewGateway_DefaultBaseURL(t *testing.T) {
	t.Parallel()

	_, err := NewGateway(apiKeyMocked, "", nil)
	assert.NoError(t, err)
}
