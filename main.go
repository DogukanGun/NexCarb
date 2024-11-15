package main

import (
	"SensorManager/env"
	"SensorManager/utils"
	"context"
)

func main() {
	ctx := context.Background()
	utils.RunWithHandlingError(env.SetEnv(ctx))
}
