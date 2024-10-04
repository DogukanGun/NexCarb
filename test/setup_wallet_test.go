package test

import (
	"SensorManager/setup_env"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetupWallet(t *testing.T) {
	wallet := setup_env.SetupWallet()
	fmt.Printf("Wallet PK: %s", wallet.PublicKey.String())
	assert.NotEqual(t, nil, wallet)
	assert.Equal(t, true, len(wallet.PublicKey.String()) > 0)
}
