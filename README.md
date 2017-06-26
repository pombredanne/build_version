# Git Build and Tag in Go programs

Example showing how to get tag and build version in a Go main program.

`go run -ldflags "-X main.BuildVersion=$(git rev-parse HEAD) -X main.Tag=$(git describe --abbrev=0 --tags)" build_verison.go`

or `go build ...` or `go install ...`
