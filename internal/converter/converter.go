package converter

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"gopkg.in/yaml.v3"
)

// Converter handles the conversion from Markdown to PDF
type Converter struct {
	config *Config
	stats  *ConversionStats
}

// New creates a new converter instance
func New(config *Config) *Converter {
	return &Converter{
		config: config,
		stats: &ConversionStats{
			StartTime: time.Now(),
		},
	}
}

// Convert performs the conversion from Markdown to PDF
func (c *Converter) Convert() error {
	// Read input file
	content, err := os.ReadFile(c.config.InputPath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	c.stats.InputSize = int64(len(content))

	// Parse front matter and content
	frontMatter, markdownContent, err := c.parseFrontMatter(content)
	if err != nil {
		return fmt.Errorf("failed to parse front matter: %w", err)
	}

	// Merge configuration with front matter
	c.mergeConfigWithFrontMatter(frontMatter)

	// Convert markdown to HTML
	htmlContent, err := c.markdownToHTML(markdownContent)
	if err != nil {
		return fmt.Errorf("failed to convert markdown to HTML: %w", err)
	}

	// Apply template and styling
	styledHTML, err := c.applyTemplate(htmlContent, frontMatter)
	if err != nil {
		return fmt.Errorf("failed to apply template: %w", err)
	}

	// Convert HTML to PDF
	err = c.htmlToPDF(styledHTML)
	if err != nil {
		return fmt.Errorf("failed to convert HTML to PDF: %w", err)
	}

	// Update stats
	c.stats.EndTime = time.Now()
	c.stats.ProcessingMS = c.stats.EndTime.Sub(c.stats.StartTime).Milliseconds()

	return nil
}

// parseFrontMatter extracts YAML front matter from markdown content
func (c *Converter) parseFrontMatter(content []byte) (*FrontMatter, []byte, error) {
	contentStr := string(content)
	frontMatter := &FrontMatter{}

	// Check if content starts with front matter
	if !strings.HasPrefix(contentStr, "---\n") {
		return frontMatter, content, nil
	}

	// Find the closing front matter delimiter
	lines := strings.Split(contentStr, "\n")
	var frontMatterLines []string
	var contentLines []string

	inFrontMatter := false
	frontMatterClosed := false

	for i, line := range lines {
		if i == 0 && line == "---" {
			inFrontMatter = true
			continue
		}

		if inFrontMatter && line == "---" {
			frontMatterClosed = true
			inFrontMatter = false
			contentLines = lines[i+1:]
			break
		}

		if inFrontMatter {
			frontMatterLines = append(frontMatterLines, line)
		}
	}

	if !frontMatterClosed {
		return frontMatter, content, nil
	}

	// Parse YAML front matter
	if len(frontMatterLines) > 0 {
		yamlContent := strings.Join(frontMatterLines, "\n")
		if err := yaml.Unmarshal([]byte(yamlContent), frontMatter); err != nil {
			return nil, nil, fmt.Errorf("invalid YAML front matter: %w", err)
		}
	}

	markdownContent := strings.Join(contentLines, "\n")
	return frontMatter, []byte(markdownContent), nil
}

// mergeConfigWithFrontMatter merges front matter settings with config
func (c *Converter) mergeConfigWithFrontMatter(fm *FrontMatter) {
	if fm.Theme != "" {
		c.config.Theme = fm.Theme
	}
	if fm.Template != "" {
		c.config.TemplateName = fm.Template
	}
	if fm.CSS != "" {
		c.config.CSSPath = fm.CSS
	}
}

// markdownToHTML converts markdown content to HTML
func (c *Converter) markdownToHTML(content []byte) (string, error) {
	// Configure goldmark with extensions
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,            // GitHub Flavored Markdown
			extension.Table,          // Tables
			extension.Strikethrough,  // Strikethrough
			extension.TaskList,       // Task lists
			extension.DefinitionList, // Definition lists
			highlighting.NewHighlighting( // Syntax highlighting
				highlighting.WithStyle("github"),
				highlighting.WithGuessLanguage(true),
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert(content, &buf); err != nil {
		return "", fmt.Errorf("markdown conversion failed: %w", err)
	}

	htmlContent := buf.String()

	// Process table of contents if requested
	htmlContent = c.processTableOfContents(htmlContent)

	return htmlContent, nil
}

// processTableOfContents generates and inserts table of contents
func (c *Converter) processTableOfContents(html string) string {
	if !strings.Contains(html, "<!-- TOC -->") {
		return html
	}

	// Extract headings
	headingRegex := regexp.MustCompile(`<h([1-6])[^>]*id="([^"]*)"[^>]*>([^<]+)</h[1-6]>`)
	matches := headingRegex.FindAllStringSubmatch(html, -1)

	if len(matches) == 0 {
		return strings.ReplaceAll(html, "<!-- TOC -->", "")
	}

	// Build TOC
	var tocBuilder strings.Builder
	tocBuilder.WriteString("<div class=\"toc\">\n<h2>Table of Contents</h2>\n<ul>\n")

	for _, match := range matches {
		level := match[1]
		id := match[2]
		title := strings.TrimSpace(match[3])

		// Create appropriate indentation
		indent := ""
		if level != "1" {
			for i := 1; i < len(level); i++ {
				indent += "  "
			}
		}

		tocBuilder.WriteString(fmt.Sprintf("%s<li><a href=\"#%s\">%s</a></li>\n", indent, id, title))
	}

	tocBuilder.WriteString("</ul>\n</div>\n")

	return strings.ReplaceAll(html, "<!-- TOC -->", tocBuilder.String())
}

// applyTemplate applies the template and styling to HTML content
func (c *Converter) applyTemplate(content string, frontMatter *FrontMatter) (string, error) {
	// Load template
	template, err := c.loadTemplate()
	if err != nil {
		return "", err
	}

	// Load CSS
	css, err := c.loadCSS()
	if err != nil {
		return "", err
	}

	// Replace template placeholders
	result := template
	result = strings.ReplaceAll(result, "{{TITLE}}", c.getTitle(frontMatter))
	result = strings.ReplaceAll(result, "{{CSS}}", css)
	result = strings.ReplaceAll(result, "{{CONTENT}}", content)

	return result, nil
}

func (c *Converter) getTitle(fm *FrontMatter) string {
	if fm.Title != "" {
		return fm.Title
	}
	return filepath.Base(c.config.InputPath)
}

// htmlToPDF converts HTML content to PDF using chromedp
func (c *Converter) htmlToPDF(htmlContent string) error {
	return c.htmlToPDFChrome(htmlContent)
}

// htmlToPDFChrome converts HTML content to PDF using Chrome headless
func (c *Converter) htmlToPDFChrome(htmlContent string) error {
	// Create a temporary HTML file
	tempDir := os.TempDir()
	tempHTMLPath := filepath.Join(tempDir, fmt.Sprintf("pdfy_%d.html", time.Now().UnixNano()))

	if err := os.WriteFile(tempHTMLPath, []byte(htmlContent), 0o644); err != nil {
		return fmt.Errorf("failed to write temporary HTML file: %w", err)
	}
	defer os.Remove(tempHTMLPath)

	// Create Chrome context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Set timeout
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var pdfBuffer []byte

	// Navigate to the HTML file and generate PDF
	err := chromedp.Run(ctx,
		chromedp.Navigate("file://"+tempHTMLPath),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Get PDF with custom options
			buf, _, err := page.PrintToPDF().
				WithPrintBackground(true).
				WithPaperWidth(8.27).  // A4 width in inches
				WithPaperHeight(11.7). // A4 height in inches
				WithMarginTop(0.4).
				WithMarginBottom(0.4).
				WithMarginLeft(0.4).
				WithMarginRight(0.4).
				WithDisplayHeaderFooter(false).
				Do(ctx)
			if err != nil {
				return err
			}

			pdfBuffer = buf
			return nil
		}),
	)
	if err != nil {
		return fmt.Errorf("failed to generate PDF: %w", err)
	}

	// Write PDF to output file
	if err := os.WriteFile(c.config.OutputPath, pdfBuffer, 0o644); err != nil {
		return fmt.Errorf("failed to write PDF file: %w", err)
	}

	// Update stats
	c.stats.OutputSize = int64(len(pdfBuffer))

	return nil
}

// GetStats returns conversion statistics
func (c *Converter) GetStats() *ConversionStats {
	return c.stats
}
