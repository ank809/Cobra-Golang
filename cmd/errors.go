package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func someFunc() error {
	return fmt.Errorf("some conditions not met")
}

var Try = &cobra.Command{
	Use:   "try",
	Short: "Try and possibly fail at something",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := someFunc(); err != nil {
			return err
		}
		return nil
	},
}
