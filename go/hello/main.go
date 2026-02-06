package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "app",
		Short: "Bazel + Cobra app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Cobra is running inside Bazel!")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
