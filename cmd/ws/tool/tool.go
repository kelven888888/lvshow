package tool

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "app [command]",
}

func Execute() {
	rootCmd.Execute()
}
