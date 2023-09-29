package uuid

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func NewCmdUuid() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uuid",
		Short: "Get a random uuid",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(uuid.NewString())
		},
	}
	return cmd
}
