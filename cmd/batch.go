package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"pdfy/internal/converter"

	"github.com/spf13/cobra"
)

var outputDir string

var batchCmd = &cobra.Command{
	Use:   "batch [glob-pattern]",
	Short: "Convert multiple Markdown files to PDF",
	Long: `Convert multiple Markdown files matching a glob pattern to PDF.

Examples:
  pdfy batch "*.md" --output-dir pdfs/
  pdfy batch "docs/**/*.md" --template technical
  pdfy batch "./markdown_files/*.md"`,
	Args: cobra.ExactArgs(1),
	RunE: batchConvert,
}

func init() {
	batchCmd.Flags().StringVar(&outputDir, "output-dir", "", "Output directory for PDF files")
	batchCmd.Flags().StringVarP(&templateName, "template", "t", "default", "Template to use")
	batchCmd.Flags().StringVar(&cssPath, "css", "", "Custom CSS file path")
	batchCmd.Flags().StringVar(&theme, "theme", "light", "Theme to use")
}

func batchConvert(cmd *cobra.Command, args []string) error {
	pattern := args[0]

	// Find matching files
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("invalid glob pattern: %w", err)
	}

	if len(matches) == 0 {
		return fmt.Errorf("no files found matching pattern: %s", pattern)
	}

	// Create output directory if specified
	if outputDir != "" {
		if err := os.MkdirAll(outputDir, 0o755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}
	}

	fmt.Printf("Found %d files to convert\n", len(matches))

	successCount := 0
	for _, inputPath := range matches {
		// Skip non-markdown files
		if !isMarkdownFile(inputPath) {
			continue
		}

		// Generate output path
		var outPath string
		if outputDir != "" {
			filename := filepath.Base(inputPath)
			ext := filepath.Ext(filename)
			pdfName := strings.TrimSuffix(filename, ext) + ".pdf"
			outPath = filepath.Join(outputDir, pdfName)
		} else {
			ext := filepath.Ext(inputPath)
			outPath = strings.TrimSuffix(inputPath, ext) + ".pdf"
		}

		// Convert file
		config := &converter.Config{
			InputPath:    inputPath,
			OutputPath:   outPath,
			TemplateName: templateName,
			CSSPath:      cssPath,
			Theme:        theme,
		}

		conv := converter.New(config)

		fmt.Printf("Converting %s...", filepath.Base(inputPath))

		if err := conv.Convert(); err != nil {
			fmt.Printf(" ✗ Failed: %v\n", err)
			continue
		}

		fmt.Printf(" ✓ Success\n")
		successCount++
	}

	fmt.Printf("\nCompleted: %d/%d files converted successfully\n", successCount, len(matches))
	return nil
}

func isMarkdownFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".md" || ext == ".markdown"
}
