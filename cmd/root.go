package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "blog",
	Short: "A little helper for managing Blog",
	Long: `BLOG v0.5
	This is a CLI program for better managing Blog
	Developed with cobra`,
}

func init() {
	RootCmd.AddCommand(initCmd)

}
