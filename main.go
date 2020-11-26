package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func numeronyze(in string) string {
	if len(in) == 0 {
		return ""
	}
	in = strings.Replace(in, " ", "", 0)
	mid := len(in) - 2
	last := string(in[len(in)-1])
	return fmt.Sprintf("%s%d%s", string(in[0]), mid, last)

}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "n9r [input]",
	Short: "numeronyzer numeronyzes strings",
	Long: `numeronyzer (or n9r) generates the numeronym for the provided input.

For example:
$ n9r numeronyzer
> n9r

or
$ n9r kubernetes
> k8s

`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
			return
		}
		in := args[0]
		resp := numeronyze(in)
		fmt.Println(resp)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
