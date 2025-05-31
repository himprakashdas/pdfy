package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/himprakashdas/pdfy/internal/converter"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch [directory]",
	Short: "Watch directory for changes and auto-convert",
	Long: `Watch a directory for Markdown file changes and automatically convert them to PDF.

Examples:
  pdfy watch docs/ --output-dir pdfs/
  pdfy watch . --template technical
  pdfy watch ./content`,
	Args: cobra.ExactArgs(1),
	RunE: watchDirectory,
}

func init() {
	watchCmd.Flags().StringVar(&outputDir, "output-dir", "", "Output directory for PDF files")
	watchCmd.Flags().StringVarP(&templateName, "template", "t", "default", "Template to use")
	watchCmd.Flags().StringVar(&cssPath, "css", "", "Custom CSS file path")
	watchCmd.Flags().StringVar(&theme, "theme", "light", "Theme to use")
}

func watchDirectory(cmd *cobra.Command, args []string) error {
	watchDir := args[0]

	// Create file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create watcher: %w", err)
	}
	defer watcher.Close()

	// Add directory to watcher
	err = watcher.Add(watchDir)
	if err != nil {
		return fmt.Errorf("failed to watch directory: %w", err)
	}

	fmt.Printf("Watching %s for changes... (Press Ctrl+C to stop)\n", watchDir)

	// Track recent conversions to avoid duplicate processing
	recentlyProcessed := make(map[string]time.Time)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}

			// Only process write events on markdown files
			if event.Op&fsnotify.Write == fsnotify.Write && isMarkdownFile(event.Name) {
				// Debounce rapid file changes
				if lastProcessed, exists := recentlyProcessed[event.Name]; exists {
					if time.Since(lastProcessed) < 2*time.Second {
						continue
					}
				}

				recentlyProcessed[event.Name] = time.Now()

				fmt.Printf("Change detected: %s\n", filepath.Base(event.Name))

				// Convert the file
				if err := convertWatchedFile(event.Name); err != nil {
					log.Printf("Conversion failed for %s: %v", event.Name, err)
				} else {
					fmt.Printf("✓ Converted %s\n", filepath.Base(event.Name))
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			log.Printf("Watcher error: %v", err)
		}
	}
}

func convertWatchedFile(inputPath string) error {
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
	return conv.Convert()
}
