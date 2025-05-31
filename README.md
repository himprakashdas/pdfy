# Pdfy 📄✨

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/Build-Passing-brightgreen.svg)](https://github.com/yourusername/pdfy)

A modern, fast, and feature-rich CLI tool for converting Markdown files to beautifully formatted PDFs. Built with Go and powered by Chrome's headless PDF engine for professional-quality output.

## ✨ Features

- 🚀 **Lightning-fast conversion** using Chrome headless engine
- 📝 **GitHub Flavored Markdown** support with tables, task lists, and more
- 🎨 **Professional themes** (light, dark) with customizable CSS
- 🌈 **Syntax highlighting** for 100+ programming languages
- 📋 **YAML front matter** for document metadata and configuration
- 📚 **Table of Contents** generation with `<!-- TOC -->` placeholder
- 🔄 **Batch processing** with glob pattern support
- 👀 **Watch mode** for real-time conversion during development
- 🎯 **Zero dependencies** - single binary deployment
- 📱 **Cross-platform** support (Windows, macOS, Linux)

## 🚀 Quick Start

### Installation

**Option 1: Download Binary** (Recommended)

```bash
# Download the latest release for your platform
curl -L https://github.com/yourusername/pdfy/releases/latest/download/pdfy-linux-amd64 -o pdfy
chmod +x pdfy
sudo mv pdfy /usr/local/bin/
```

**Option 2: Build from Source**

```bash
git clone https://github.com/yourusername/pdfy.git
cd pdfy
go build -o pdfy
```

**Option 3: Go Install**

```bash
go install github.com/yourusername/pdfy@latest
```

### Basic Usage

```bash
# Convert a single file
pdfy convert README.md

# Specify output location
pdfy convert docs/guide.md --output guides/guide.pdf

# Use dark theme
pdfy convert document.md --theme dark
```

## 📖 Comprehensive Guide

### Single File Conversion

```bash
# Basic conversion
pdfy convert document.md

# With custom output path
pdfy convert document.md -o /path/to/output.pdf

# Using different theme
pdfy convert document.md --theme dark

# With custom CSS
pdfy convert document.md --css custom-styles.css
```

### Batch Processing

Perfect for converting multiple files at once:

```bash
# Convert all markdown files in current directory
pdfy batch "*.md"

# Convert files in subdirectories
pdfy batch "docs/**/*.md" --output-dir pdfs/

# Convert with specific pattern
pdfy batch "chapter-*.md" -o book-chapters/
```

### Watch Mode

Ideal for development workflows:

```bash
# Watch current directory
pdfy watch .

# Watch specific directory with output folder
pdfy watch docs/ --output-dir build/pdfs/

# Watch with specific theme
pdfy watch . --theme dark
```

### YAML Front Matter

Enhance your documents with metadata:

```yaml
---
title: "Project Documentation"
theme: dark
css: custom.css
---
# Your Markdown Content

Your document content goes here...
```

### Table of Contents

Add `<!-- TOC -->` anywhere in your markdown to generate an automatic table of contents:

```markdown
# Document Title

<!-- TOC -->

## Section 1

Content here...

## Section 2

Content here...
```

## 🎨 Themes & Customization

### Built-in Themes

- **light** (default) - Clean, professional appearance
- **dark** - Dark mode with syntax highlighting

### Custom CSS

```bash
pdfy convert document.md --css my-styles.css
```

Or specify in YAML front matter:

```yaml
---
css: path/to/custom.css
---
```

## 📋 Examples

### Technical Documentation

```bash
# Convert API documentation
pdfy convert api-docs.md --theme light -o api-reference.pdf
```

Input (`api-docs.md`):

````markdown
---
title: "API Reference Guide"
---

# API Documentation

<!-- TOC -->

## Authentication

```bash
curl -H "Authorization: Bearer TOKEN" https://api.example.com
```
````

## Endpoints

| Method | Endpoint | Description |
| ------ | -------- | ----------- |
| GET    | /users   | List users  |
| POST   | /users   | Create user |

````

### Book/Report Generation

```bash
# Batch convert book chapters
pdfy batch "chapter-*.md" --output-dir book/
````

## 🛠️ Advanced Configuration

### Custom Templates

Create your own HTML templates in `internal/converter/templates/`:

```html
<!DOCTYPE html>
<html>
  <head>
    <title>{{TITLE}}</title>
    <style>
      {{CSS}}
    </style>
  </head>
  <body>
    <div class="document">{{CONTENT}}</div>
  </body>
</html>
```

### Environment Variables

```bash
export PDFY_THEME=dark
export PDFY_OUTPUT_DIR=./pdfs
```

## 🔧 Requirements

- **Go 1.21+** (for building from source)
- **Chrome/Chromium** (for PDF generation)
- **Linux/macOS/Windows** (cross-platform support)

## 🐛 Troubleshooting

### Common Issues

**Chrome not found:**

```bash
# Install Chrome/Chromium
# Ubuntu/Debian
sudo apt install chromium-browser

# macOS
brew install --cask google-chrome

# Or set custom Chrome path
export CHROME_BIN=/path/to/chrome
```

**Permission denied:**

```bash
chmod +x pdfy
```

**Large files timing out:**

```bash
# Files are processed with 30s timeout by default
# For very large files, consider splitting them
```

## 📊 Performance

- **Small files** (<1MB): ~100-500ms
- **Medium files** (1-10MB): ~500ms-2s
- **Large files** (10MB+): ~2-10s

Batch processing is parallelized for optimal performance.

## 🤝 Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

- 🐛 **Bug reports** via GitHub Issues
- 💡 **Feature requests** via GitHub Discussions
- 🔧 **Pull requests** for bug fixes and features
- 📖 **Documentation** improvements

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Goldmark](https://github.com/yuin/goldmark) - Markdown parser
- [Chroma](https://github.com/alecthomas/chroma) - Syntax highlighting
- [ChromeDP](https://github.com/chromedp/chromedp) - Chrome automation
- [Cobra](https://github.com/spf13/cobra) - CLI framework

## 📞 Support

- 📚 **Documentation**: [Wiki](https://github.com/yourusername/pdfy/wiki)
- 💬 **Discussions**: [GitHub Discussions](https://github.com/yourusername/pdfy/discussions)
- 🐛 **Issues**: [GitHub Issues](https://github.com/yourusername/pdfy/issues)

---

**Made with ❤️ in Go** | **Star ⭐ if you found this helpful!**
