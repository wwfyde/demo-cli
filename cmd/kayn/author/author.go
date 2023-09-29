package author

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	name  = "wwfyde"
	email = "wwfyde@gmail.com"
)

func NewCmdAuthor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "author",
		Short: "Get the author info ",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Name:", name)
			fmt.Println("Email:", email)
		},
	}
	return cmd
}
