package packb

import (
	"time"
)

func Sleep1SecAndEchoTwice(input string) string {
	time.Sleep(1 * time.Second)
	return input + input
}
