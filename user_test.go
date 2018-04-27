package dd

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetUserInfo(t *testing.T) {
	_, err := client.GetUserInfo("1")
	assert.Error(t, err)
}
