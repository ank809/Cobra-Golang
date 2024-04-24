package cmd

/* Persistent Flags
A flag can be 'persistent', meaning that this flag will be available to the command it's assigned to as well as every command under that command.
For global flags, assign a flag as a persistent flag on the root.

Local Flags
A flag can also be assigned locally, which will only apply to that specific command.

*/
import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version123",
	Short: "The version of this todo app is 1.1.1",
	Long:  "Version can be upadted later on ",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("User --help for more info")
			return
		}
	},
}

var s string

func init() {
	VersionCmd.PersistentFlags().StringVarP(&s, "user", "u", "", "username(required)")
	VersionCmd.MarkPersistentFlagRequired("user")
}
