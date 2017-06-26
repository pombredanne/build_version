package main

import "fmt"
import "runtime"

var (
	// BuildVersion gets set to the current git SHA
	BuildVersion = "unknown"
	// Tag also gets set to the most recent tag
	Tag = "unknown"
)

func main() {
	fmt.Printf("Hello! Latest tag is %s, Build version is %s, built for %s/%s\n", Tag, BuildVersion, runtime.GOOS, runtime.GOARCH)
}
