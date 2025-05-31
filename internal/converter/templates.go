package converter

import (
	"embed"
	"fmt"
	"os"
	"strings"
)

//go:embed templates/*
var templatesFS embed.FS

//go:embed themes/*
var themesFS embed.FS

// loadTemplate loads the HTML template based on the configuration
func (c *Converter) loadTemplate() (string, error) {
	templateName := c.config.TemplateName
	if templateName == "" {
		templateName = "default"
	}

	// Try to load from embedded templates first
	templatePath := fmt.Sprintf("templates/%s.html", templateName)
	if content, err := templatesFS.ReadFile(templatePath); err == nil {
		return string(content), nil
	}

	// Fall back to basic template
	return c.getDefaultTemplate(), nil
}

// loadCSS loads CSS styles based on theme and custom CSS
func (c *Converter) loadCSS() (string, error) {
	var cssBuilder strings.Builder

	// Load theme CSS
	theme := c.config.Theme
	if theme == "" {
		theme = "light"
	}

	themePath := fmt.Sprintf("themes/%s.css", theme)
	if themeCSS, err := themesFS.ReadFile(themePath); err == nil {
		cssBuilder.Write(themeCSS)
		cssBuilder.WriteString("\n")
	} else {
		// Fall back to default styles
		cssBuilder.WriteString(c.getDefaultCSS())
		cssBuilder.WriteString("\n")
	}

	// Load custom CSS if provided
	if c.config.CSSPath != "" {
		customCSS, err := os.ReadFile(c.config.CSSPath)
		if err != nil {
			return "", fmt.Errorf("failed to read custom CSS file: %w", err)
		}
		cssBuilder.Write(customCSS)
	}

	return cssBuilder.String(), nil
}

// getDefaultTemplate returns a basic HTML template
func (c *Converter) getDefaultTemplate() string {
	return `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{TITLE}}</title>
    <meta name="author" content="{{AUTHOR}}">
    <style>
        {{CSS}}
    </style>
</head>
<body>
    <div class="document">
        <header class="page-header">
            {{HEADER}}
        </header>
        
        <main class="content">
            <div class="title-page">
                <h1 class="document-title">{{TITLE}}</h1>
                <p class="document-author">{{AUTHOR}}</p>
                <p class="document-date">{{DATE}}</p>
            </div>
            
            <div class="document-content">
                {{CONTENT}}
            </div>
        </main>
        
        <footer class="page-footer">
            {{FOOTER}}
        </footer>
    </div>
</body>
</html>`
}

// getDefaultCSS returns basic CSS styles
func (c *Converter) getDefaultCSS() string {
	return `
/* Base styles */
* {
    box-sizing: border-box;
}

body {
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', sans-serif;
    line-height: 1.6;
    color: #333;
    margin: 0;
    padding: 20px;
    background: white;
}

.document {
    max-width: 210mm;
    margin: 0 auto;
    background: white;
    box-shadow: 0 0 10px rgba(0,0,0,0.1);
    padding: 40px;
}

/* Typography */
h1, h2, h3, h4, h5, h6 {
    margin-top: 2em;
    margin-bottom: 1em;
    line-height: 1.2;
    page-break-after: avoid;
}

h1 { font-size: 2.5em; color: #2c3e50; }
h2 { font-size: 2em; color: #34495e; border-bottom: 2px solid #eee; padding-bottom: 0.3em; }
h3 { font-size: 1.5em; color: #34495e; }
h4 { font-size: 1.2em; }
h5 { font-size: 1.1em; }
h6 { font-size: 1em; }

p {
    margin-bottom: 1em;
    text-align: justify;
}

/* Code blocks */
code {
    font-family: 'SFMono-Regular', 'Consolas', 'Liberation Mono', 'Menlo', monospace;
    background-color: #f8f9fa;
    padding: 2px 6px;
    border-radius: 3px;
    font-size: 0.9em;
}

pre {
    background-color: #f8f9fa;
    border: 1px solid #e9ecef;
    border-radius: 6px;
    padding: 16px;
    overflow-x: auto;
    margin: 1em 0;
    page-break-inside: avoid;
}

pre code {
    background: none;
    padding: 0;
    border-radius: 0;
}

/* Tables */
table {
    width: 100%;
    border-collapse: collapse;
    margin: 1em 0;
    page-break-inside: avoid;
}

th, td {
    border: 1px solid #ddd;
    padding: 12px;
    text-align: left;
}

th {
    background-color: #f8f9fa;
    font-weight: 600;
}

tr:nth-child(even) {
    background-color: #f8f9fa;
}

/* Lists */
ul, ol {
    margin: 1em 0;
    padding-left: 2em;
}

li {
    margin-bottom: 0.5em;
}

/* Blockquotes */
blockquote {
    margin: 1em 0;
    padding: 0 1em;
    border-left: 4px solid #3498db;
    background-color: #f8f9fa;
    font-style: italic;
}

/* Links */
a {
    color: #3498db;
    text-decoration: none;
}

a:hover {
    text-decoration: underline;
}

/* Images */
img {
    max-width: 100%;
    height: auto;
    display: block;
    margin: 1em auto;
}

/* Table of Contents */
.toc {
    background-color: #f8f9fa;
    border: 1px solid #e9ecef;
    border-radius: 6px;
    padding: 20px;
    margin: 2em 0;
    page-break-inside: avoid;
}

.toc h2 {
    margin-top: 0;
    margin-bottom: 1em;
    border-bottom: none;
}

.toc ul {
    list-style: none;
    padding-left: 0;
}

.toc li {
    margin-bottom: 0.3em;
}

.toc a {
    text-decoration: none;
    color: #333;
}

.toc a:hover {
    color: #3498db;
}

/* Title page */
.title-page {
    text-align: center;
    margin-bottom: 3em;
    page-break-after: always;
}

.document-title {
    font-size: 3em;
    margin-bottom: 0.5em;
    color: #2c3e50;
}

.document-author {
    font-size: 1.2em;
    color: #7f8c8d;
    margin-bottom: 0.5em;
}

.document-date {
    color: #95a5a6;
}

/* Headers and footers */
.page-header, .page-footer {
    font-size: 0.9em;
    color: #7f8c8d;
    text-align: center;
    padding: 10px 0;
}

.page-header {
    border-bottom: 1px solid #eee;
    margin-bottom: 2em;
}

.page-footer {
    border-top: 1px solid #eee;
    margin-top: 2em;
}

/* Print styles */
@media print {
    body {
        padding: 0;
    }
    
    .document {
        box-shadow: none;
        max-width: none;
        padding: 0;
    }
    
    /* Page breaks */
    h1, h2, h3, h4, h5, h6 {
        page-break-after: avoid;
    }
    
    pre, blockquote, table, .toc {
        page-break-inside: avoid;
    }
    
    /* Hide elements that shouldn't print */
    .no-print {
        display: none !important;
    }
}

/* Page numbering for PDF */
@page {
    margin: 2.5cm;
    
    @bottom-center {
        content: counter(page) " / " counter(pages);
        font-size: 9pt;
        color: #666;
    }
}

/* Syntax highlighting styles */
.chroma .err { color: #a61717; background-color: #e3d2d2 }
.chroma .k { color: #000000; font-weight: bold }
.chroma .ch { color: #999988; font-style: italic }
.chroma .cm { color: #999988; font-style: italic }
.chroma .cp { color: #999999; font-weight: bold }
.chroma .cpf { color: #999988; font-style: italic }
.chroma .c1 { color: #999988; font-style: italic }
.chroma .cs { color: #999999; font-weight: bold; font-style: italic }
.chroma .gd { color: #000000; background-color: #ffdddd }
.chroma .ge { color: #000000; font-style: italic }
.chroma .gr { color: #aa0000 }
.chroma .gh { color: #999999 }
.chroma .gi { color: #000000; background-color: #ddffdd }
.chroma .go { color: #888888 }
.chroma .gp { color: #555555 }
.chroma .gs { font-weight: bold }
.chroma .gu { color: #aaaaaa }
.chroma .gt { color: #aa0000 }
.chroma .kc { color: #000000; font-weight: bold }
.chroma .kd { color: #000000; font-weight: bold }
.chroma .kn { color: #000000; font-weight: bold }
.chroma .kp { color: #000000; font-weight: bold }
.chroma .kr { color: #000000; font-weight: bold }
.chroma .kt { color: #445588; font-weight: bold }
.chroma .m { color: #009999 }
.chroma .s { color: #d01040 }
.chroma .na { color: #008080 }
.chroma .nb { color: #0086B3 }
.chroma .nc { color: #445588; font-weight: bold }
.chroma .no { color: #008080 }
.chroma .nd { color: #3c5d5d; font-weight: bold }
.chroma .ni { color: #800080 }
.chroma .ne { color: #990000; font-weight: bold }
.chroma .nf { color: #990000; font-weight: bold }
.chroma .nl { color: #990000; font-weight: bold }
.chroma .nn { color: #555555 }
.chroma .nt { color: #000080 }
.chroma .nv { color: #008080 }
.chroma .ow { color: #000000; font-weight: bold }
.chroma .w { color: #bbbbbb }
.chroma .mb { color: #009999 }
.chroma .mf { color: #009999 }
.chroma .mh { color: #009999 }
.chroma .mi { color: #009999 }
.chroma .mo { color: #009999 }
.chroma .sa { color: #d01040 }
.chroma .sb { color: #d01040 }
.chroma .sc { color: #d01040 }
.chroma .dl { color: #d01040 }
.chroma .sd { color: #d01040 }
.chroma .s2 { color: #d01040 }
.chroma .se { color: #d01040 }
.chroma .sh { color: #d01040 }
.chroma .si { color: #d01040 }
.chroma .sx { color: #d01040 }
.chroma .sr { color: #009926 }
.chroma .s1 { color: #d01040 }
.chroma .ss { color: #990073 }
.chroma .bp { color: #999999 }
.chroma .fm { color: #990000; font-weight: bold }
.chroma .vc { color: #008080 }
.chroma .vg { color: #008080 }
.chroma .vi { color: #008080 }
.chroma .vm { color: #008080 }
.chroma .il { color: #009999 }
`
}
