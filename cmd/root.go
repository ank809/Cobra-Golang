package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var short_text = "This is a simple CLI based Todo Application in golang"
var long_text = "You can \n create \n add \n read \n delete todos "

var rootCmd = &cobra.Command{
	Use:   "Todo List",
	Short: short_text,
	Long:  long_text,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(short_text)
			fmt.Println(long_text)
			return
		}
	},
}
