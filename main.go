package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hookenz/gotailwind/v4/downloader"
)

func main() {
	tailwindPath, err := downloader.EnsureTailwindInstalled("v4.1.5")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to install tailwind: %v\n", err)
		os.Exit(1)
	}

	cmd := exec.Command(tailwindPath, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		os.Exit(cmd.ProcessState.ExitCode())
	}
}
