package main

import "fmt"
import "runtime"

var (
	// BuildVersion gets set to the current git SHA while building
	BuildVersion = "unknown"
)

func main() {
	fmt.Printf("Hello! Build version is %s, running on %s/%s\n", BuildVersion, runtime.GOOS, runtime.GOARCH)
}
