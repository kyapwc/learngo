package cmd

import (
	"fmt"
	"runtime"
)

func Version() {
	fmt.Println(runtime.Version())
}
