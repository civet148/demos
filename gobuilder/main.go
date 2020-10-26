package main

import (
	"fmt"
)

var (
	AppName     string
	AppVersion  string
	GitRevision string
	GitBranch   string
	GoVersion   string
	BuildTime   string
)

func main() {
	fmt.Printf("App Name:\t%s\n", AppName)
	fmt.Printf("App Version:\t%s\n", AppVersion)
	fmt.Printf("Build time:\t%s\n", BuildTime)
	fmt.Printf("Git revision:\t%s\n", GitRevision)
	fmt.Printf("Git branch:\t%s\n", GitBranch)
	fmt.Printf("Golang Version: %s\n", GoVersion)
}
