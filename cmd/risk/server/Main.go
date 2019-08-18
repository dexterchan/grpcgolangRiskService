package main

import (
	"fmt"
	"os"

	cmd "github.com/dexter/grpcRiskStandalone/pkg/cmd"
)

func main() {
	if err := cmd.RunRiskServerCmdLine(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
