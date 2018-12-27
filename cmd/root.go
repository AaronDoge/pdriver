package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const CURRENT_VERSION="v0.1.0"

var rootCmd = &cobra.Command{
	Use: 	"pdriver",
	Short: 	"pdriver is process driver. ",
	Long: 	`A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if Version {
			fmt.Printf("pdriver %s.\n", CURRENT_VERSION)
			return nil
		}
		return cmd.Usage()
	},
}

var (
	Version 	bool
	Namespace	string
	Token 	 	string
	Region 		string
	Filename 	string
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("root cmd execute error. ", err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(CheckCmd)

	initUploadCmd()

	rootCmd.Flags().BoolVarP(&Version, "version","v", false, "show version")
	UploadCmd.PersistentFlags().StringVarP(&Namespace, "namespace", "n", "", "Valid namespace should be specified.")
	UploadCmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "Valid token should be specified.")
	UploadCmd.PersistentFlags().StringVarP(&Region, "region", "r", "", "Valid token should be region.")

}
