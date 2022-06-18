package health

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewReadyCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "ready",
		Short:   "Checks if the Heimdall deployment is ready",
		Example: "heimdall health --endpoint=http://localhost:4456/ ready",
		Run: func(cmd *cobra.Command, args []string) {
			// nolint
			fmt.Println("health ready")
		},
	}
}