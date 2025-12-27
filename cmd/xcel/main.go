package main

import (
	"fmt"
	"os"

	"xcel/internal/commands"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xcel",
	Short: "XCel is a powerful Excel processing framework",
	Long:  `XCel is a powerful Excel processing framework similar to MSFConsole.`,
}

func main() {
	// 添加子命令
	rootCmd.AddCommand(commands.NewConvertCommand())
	rootCmd.AddCommand(commands.NewColStatCommand())
	rootCmd.AddCommand(commands.NewHeadCommand())
	rootCmd.AddCommand(commands.NewTailCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
