package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

//

var CheckCmd = &cobra.Command{
	Use: 	"check",
	Short: 	"check the state",
	Long: 	`All software has versions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check the status of procedure.")

	},
}


