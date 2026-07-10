package main

import (
	"flag"
	"fmt"
	"ginshop.com/cmd/ws/tool"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// https://blog.csdn.net/m0_73651896/article/details/131076584
func main() {
	tool.Execute()
}

func GenWebCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "web",
		Short:         "Generate model, cache, dao, handler, web code",
		Long:          "generate model, cache, dao, handler, web code.",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(time.Now())
		},
	}

	return cmd
}

func rangeSlice() {
	for k, v := range os.Args {
		fmt.Println(k, v)
	}
}

// go run .\main.go -name 54444444444444
func flagpring() {
	var name string
	flag.StringVar(&name, "name", "getname", "names")
	flag.Parse()
	fmt.Println(name)
}
