package main

import (
	"SensorManager/env"
	"SensorManager/utils"
	"context"
	"os"
)

func main() {
	ctx := context.Background()
	wallet := env.SetupWallet()
	utils.RunWithHandlingError(env.SetEnv(ctx))
	utils.RunWithHandlingError(os.Setenv(env.PUBKEY, wallet.PublicKey.String()))
}
