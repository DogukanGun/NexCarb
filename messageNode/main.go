package main

import (
	env2 "SensorManager/messageNode/env"
	"SensorManager/utils"
	"context"
	"os"
)

func main() {
	ctx := context.Background()
	wallet := env2.SetupWallet()
	utils.RunWithHandlingError(env2.SetEnv(ctx))
	utils.RunWithHandlingError(os.Setenv(env2.PUBKEY, wallet.PublicKey.String()))

}
