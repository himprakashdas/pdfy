# Pdfy - Markdown to PDF Converter

<img src="https://r2cdn.perplexity.ai/pplx-full-logo-primary-dark%402x.png" width="120"/>

A powerful Go CLI tool for converting Markdown files to professionally formatted PDFs with advanced features like syntax highlighting, templates, and YAML front-matter configuration.

## Features

- ✅ **Markdown to PDF conversion** with GitHub Flavored Markdown support
- ✅ **Syntax highlighting** for code blocks
- ✅ **YAML front matter** for document configuration
- ✅ **Multiple themes** (light, dark, technical)
- ✅ **Custom CSS** support
- ✅ **Table of Contents** generation
- ✅ **Batch processing** for multiple files
- ✅ **Watch mode** for automatic conversion
- ✅ **Professional templates** for different document types
- 🔄 **PDF generation** (currently outputs HTML, PDF coming soon)

## Quick Start

### Basic Conversion

```bash
# Convert a single markdown file
pdfy convert document.md

# Specify output file
pdfy convert document.md -o output.pdf

# Use a specific theme
pdfy convert document.md --theme dark
```

### Batch Processing

```bash
# Convert all markdown files in current directory
pdfy batch "*.md"

# Convert with output directory
pdfy batch "docs/*.md" --output-dir pdfs/
```

### Watch Mode

```bash
# Watch current directory for changes
pdfy watch .

# Watch with output directory
pdfy watch docs/ --output-dir pdfs/
```

## Installation

```bash
# Build from source
go build -o pdfy main.go

# Or install directly
go install github.com/yourusername/pdfy@latest
```
