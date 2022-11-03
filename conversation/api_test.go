package conversation

import (
	"github.com/Bijles-aan-Huis-B-V/messagebird-go-rest-api/internal/mbtest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestMain(m *testing.M) {
	mbtest.EnableServer(m)
}

func TestRequestSandboxEnabled(t *testing.T) {
	data := struct{}{}
	method := http.MethodGet
	reqPath := "qwerty"

	client := mbtest.MockClient().(*mbtest.ClientMock)
	client.On("Request", data, method, apiRoot+"/"+reqPath, data).Return(nil)

	err := request(client, data, method, reqPath, data)

	assert.NoError(t, err)
}
