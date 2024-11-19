package main

import (
	env2 "SensorManager/common/env"
	"SensorManager/common/utils"
	"SensorManager/searchNode/connector"
	"context"
	"os"
)

func main() {
	ctx := context.Background()
	wallet := env2.SetupWallet()
	utils.RunWithHandlingError(env2.SetEnv(ctx, "/Users/dogukangundogan/Desktop/SensorManager/searchNode/.env"))
	utils.RunWithHandlingError(os.Setenv(env2.PUBKEY, wallet.PublicKey.String()))
	connector.Run()
}
