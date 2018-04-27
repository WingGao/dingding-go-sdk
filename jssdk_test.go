package dd

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetJSTicket(t *testing.T) {
	token, err := client.GetJSTicket()
	assert.NoError(t, err, "get ticket error")
	assert.NotEmpty(t, token.Ticket, "should get ticket")
	t.Logf("get ticket %#v", token)
}

func TestClient_GetJSSign(t *testing.T) {
	sign := client.GetJSSign("http://a.com", "1", "")
	t.Logf("js sign = %#v", sign)
}
