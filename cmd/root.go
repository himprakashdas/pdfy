package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdfy",
	Short: "A powerful Markdown to PDF converter",
	Long: `Pdfy is a CLI tool for converting Markdown files to professionally formatted PDFs.
It supports advanced features like syntax highlighting, templates, and YAML front-matter configuration.`,
	Version: "1.0.0",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(convertCmd)
	rootCmd.AddCommand(batchCmd)
	rootCmd.AddCommand(watchCmd)
}
