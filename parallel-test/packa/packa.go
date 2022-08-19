package packa

import (
	"time"
)

func Sleep1SecAndEcho(input string) string {
	time.Sleep(1 * time.Second)
	return input
}
