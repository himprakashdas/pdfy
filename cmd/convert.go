package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"pdfy/internal/converter"

	"github.com/spf13/cobra"
)

var (
	outputPath   string
	templateName string
	cssPath      string
	theme        string
)

var convertCmd = &cobra.Command{
	Use:   "convert [input.md]",
	Short: "Convert a Markdown file to PDF",
	Long: `Convert a single Markdown file to PDF with optional customization options.

Examples:
  pdfy convert document.md -o output.pdf
  pdfy convert document.md --template technical
  pdfy convert document.md --css custom.css --theme dark`,
	Args: cobra.ExactArgs(1),
	RunE: convertFile,
}

func init() {
	convertCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output PDF file path")
	convertCmd.Flags().StringVarP(&templateName, "template", "t", "default", "Template to use (default, technical, dark)")
	convertCmd.Flags().StringVar(&cssPath, "css", "", "Custom CSS file path")
	convertCmd.Flags().StringVar(&theme, "theme", "light", "Theme to use (light, dark)")
}

func convertFile(cmd *cobra.Command, args []string) error {
	inputPath := args[0]

	// Validate input file exists
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist: %s", inputPath)
	}

	// Generate output path if not provided
	if outputPath == "" {
		ext := filepath.Ext(inputPath)
		outputPath = strings.TrimSuffix(inputPath, ext) + ".pdf"
	}

	// Create converter instance
	config := &converter.Config{
		InputPath:    inputPath,
		OutputPath:   outputPath,
		TemplateName: templateName,
		CSSPath:      cssPath,
		Theme:        theme,
	}

	conv := converter.New(config)

	fmt.Printf("Converting %s to %s...\n", inputPath, outputPath)

	if err := conv.Convert(); err != nil {
		return fmt.Errorf("conversion failed: %w", err)
	}

	fmt.Printf("✓ Successfully converted to %s\n", outputPath)
	return nil
}
