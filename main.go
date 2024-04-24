package main

import (
	"fmt"
	"os"

	"github.com/ank809/Cobra-Golang/cmd"
)

func main() {

	cmd.RootCmd.AddCommand(cmd.VersionCmd)
	cmd.RootCmd.AddCommand(cmd.Try)
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
