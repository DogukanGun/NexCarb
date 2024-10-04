package main

import (
	"SensorManager/utils"
	"context"
)

func main() {
	ctx := context.Background()
	utils.RunWithHandlingError(SetEnv(ctx))
}
