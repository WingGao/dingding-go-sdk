package dd

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"fmt"
)

var (
	client *Client
)

func TestMain(m *testing.M) {
	client, _ = NewClient(os.Getenv("TEST_CORPID"), os.Getenv("TEST_CORPSECRET"))
	client.cacheToken.AccessToken = os.Getenv("TEST_DD_TOKEN")
	if client.cacheToken.AccessToken == "" {
		client.GetToken()
	}
	if client.cacheJSTicket.Ticket == "" {
		client.GetJSTicket()
	}
	fmt.Printf(" token: %s, ticket: %s\n", client.cacheToken.AccessToken, client.cacheJSTicket.Ticket)
	os.Exit(m.Run())
}

func TestNewClient(t *testing.T) {
	token, err := client.GetToken()
	assert.NoError(t, err, "get token error")
	assert.NotEmpty(t, token.AccessToken, "should get token")
	t.Logf("get token %#v", token)
}
