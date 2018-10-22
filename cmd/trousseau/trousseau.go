package main

import "github.com/spf13/cobra"

func main() {
	rootCmd := &cobra.Command{Use: "trousseau <command>"}
	rootCmd.Execute()
}
