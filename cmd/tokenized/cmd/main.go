package cmd

import (
	"github.com/spf13/cobra"
)

const (
	FlagDebugMode = "debug"
)

var scCmd = &cobra.Command{
	Use:   "tokenized",
	Short: "Tokenized CLI",
}

func Execute() {
	scCmd.AddCommand(cmdDecode)
	scCmd.AddCommand(cmdGenerateKey)
	scCmd.AddCommand(cmdGenerateExtendedKey)
	scCmd.AddCommand(cmdCreateContract)
	scCmd.Execute()
}
