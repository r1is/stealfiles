package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var zipCmd = &cobra.Command{
	Use:   "zip",
	Short: "compression module",
	Long:  "compression  file or directory",
	Run: func(cmd *cobra.Command, args []string) {
		if sourceFileName == "" || destFileName == "" {
			cmd.Help()
			return
		} else {
			fmt.Println("source file name is:", sourceFileName)
			fmt.Println("dest file name is:", destFileName)
		}
	},
}

// 需要压缩文件的名字
var (
	sourceFileName string
	destFileName   string
	allCompressor  bool
)

func init() {
	rootCmd.AddCommand(zipCmd)
	zipCmd.Flags().StringVarP(&sourceFileName, "source", "s", "", "source file name")
	zipCmd.Flags().StringVarP(&destFileName, "dest", "d", "", "source file name")
	zipCmd.Flags().BoolVarP(&allCompressor, "all", "a", false, "压缩目标下的所有文件")
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
