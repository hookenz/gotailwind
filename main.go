package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hookenz/gotailwind/v4/downloader"
)

const TailwindCssVersion = "v4.1.7"
const TailwindExtraVersion = "v2.1.24"

type config struct {
	useTailwindExtra bool
}

func main() {
	if errCode := run(config{useTailwindExtra: os.Getenv("TWCLI_EXTRA") != ""}); errCode != 0 {
		os.Exit(errCode)
	}
}

func run(cli config) int {
	var binPath string

	if cli.useTailwindExtra {
		tailwindExtraPath, err := downloader.EnsureTailwindExtraInstalled(TailwindExtraVersion)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to install tailwind-cli-extra: %v\n", err)
			os.Exit(1)
		}
		binPath = tailwindExtraPath
	} else {
		tailwindPath, err := downloader.EnsureTailwindInstalled(TailwindCssVersion)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to install tailwind: %v\n", err)
			os.Exit(1)
		}
		binPath = tailwindPath
	}

	cmd := exec.Command(binPath, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to install: %v\n", err)
		return cmd.ProcessState.ExitCode()
	}

	return 0
}
