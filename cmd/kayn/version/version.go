package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

const Version = "0.1.0"

func NewCmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Kayn",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("kayn version:", Version, "\nsee also `kayn --version`.")
		},
	}
	return cmd
}
