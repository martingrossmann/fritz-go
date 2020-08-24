package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "Fritz-Go",
	Short: "Read online counter from Fritz.box",
	Long:  "Read online counter from Fritz.box",
	Run: func(cmd *cobra.Command, args []string) {
		//fritz.Connect()
	},
}

func Exec() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
