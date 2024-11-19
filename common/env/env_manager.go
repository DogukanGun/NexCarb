package env

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func SetEnv(ctx context.Context, envPath string) (err error) {
	content, err := ioutil.ReadFile(envPath)
	if err != nil {
		log.Fatal(err)
	}
	contentByLine := strings.Split(string(content), "\n")
	for _, line := range contentByLine {
		pair := strings.Split(line, "=")
		if len(pair) != 2 {
			err = errors.New("unsupported environment file. File should be in formmat of key=value")
		}
		if pair[0] == "env" {
			ctx = context.WithValue(ctx, pair, pair[1])
		} else {
			err = os.Setenv(pair[0], pair[1])
		}
	}
	return
}
