package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "A brief description of go-cli-aviation",
	Long:  "go-cli-aviation is a CLI tool for retrieving flight information from aviationstack",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`go-cli-aviation is a CLI tool for retrieving flight information from aviationstack

      list of commands:
      about
      `)
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
