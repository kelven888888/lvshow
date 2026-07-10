package tool

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

// https://blog.csdn.net/m0_73651896/article/details/131076584
var echoTimes int
var cmdTimes = &cobra.Command{
	Use:   "time [# times] [string to echo]",
	Short: "Echo anything to the screen more times",
	Long: `echo things multiple times back to the user y providing
		a count and a string.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < echoTimes; i++ {
			fmt.Println("Echo: " + strings.Join(args, " "))
		}
	},
}

func init() {
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	cmdEcho.AddCommand(cmdTimes)
}
