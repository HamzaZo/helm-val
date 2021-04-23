package main

import (
	"github.com/HamzaZo/helm-val/cmd"
	"os"
)

func main() {
	f := cmd.NewRootCmd(os.Stdout, os.Args[1:])
	if err := f.Execute(); err != nil {
		os.Exit(1)
	}
}
