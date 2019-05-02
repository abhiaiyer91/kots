package main

import (
	"fmt"
	"os"

	"github.com/replicatedhq/ship-cluster/worker/pkg/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Println("Main process exited:", err)
		os.Exit(1)
	}
}
