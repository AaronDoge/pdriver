package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UploadCmd = &cobra.Command{
	Use: 	"upload",
	Short: 	"upload files to ufile",
	Long:	`upload files to ufile`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executing upload...")
		fmt.Println("namespace is", Namespace)
		fmt.Println("token is", Token)
	},
}

var FileName string

func initUploadCmd() {

	rootCmd.AddCommand(UploadCmd)

	UploadCmd.Flags().StringVarP(&FileName, "filename", "f", "", "file name should be specified.")

	UploadCmd.MarkFlagRequired("namespace")
	UploadCmd.MarkFlagRequired("token")
	UploadCmd.MarkFlagRequired("region")
	UploadCmd.MarkFlagRequired("filename")
}


