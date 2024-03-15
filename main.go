package main

import (
	"fmt"
	"os"

	"github.com/intility/cwc/cmd"
	"github.com/intility/cwc/pkg/ui"
)

//go:generate ./bin/lang-gen

func main() {
	command := cmd.CreateRootCommand()

	err := command.Execute()
	if err != nil {
		ui.PrintMessage(fmt.Sprintf("Error: %s\n", err), ui.MessageTypeError)
		os.Exit(1)
	}
}
