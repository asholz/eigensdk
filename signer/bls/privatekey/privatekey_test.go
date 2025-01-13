package privatekey

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOperatorId(t *testing.T) {
	signer, err := New(Config{
		PrivateKey: "13718104057011380243349384062412322292853638010146548074368241565852610884213",
	})
	assert.NoError(t, err)

	operatorId, err := signer.GetOperatorId()
	assert.NoError(t, err)
	assert.Equal(t, operatorId, "0x2f6d54c4ef253d4ac1b6bfa672a8f449a0aa6ab1c20ee97a74f8e521fe57604b")
}
