---
title: "Pdfy Example Document"
theme: light
---

# Pdfy Example Document

This is a demonstration of **Pdfy's** markdown to PDF conversion capabilities.

<!-- TOC -->

## Features Showcase

### Text Formatting

- **Bold text**
- _Italic text_
- `inline code`
- ~~strikethrough text~~

### Code Blocks

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Pdfy!")
}
```

```bash
# Convert this file to PDF
pdfy convert example.md --output example.pdf
```

### Tables

| Feature             | Status | Notes                    |
| ------------------- | ------ | ------------------------ |
| Markdown Support    | ✅     | GitHub Flavored Markdown |
| Syntax Highlighting | ✅     | 100+ languages           |
| Themes              | ✅     | Light and Dark themes    |
| Custom CSS          | ✅     | Full customization       |

### Lists

1. **Ordered lists**

   - Nested items work
   - Multiple levels supported

2. **Task lists**
   - [x] Completed task
   - [ ] Pending task
   - [x] Another completed task

### Blockquotes

> This is a blockquote that demonstrates how quoted text appears in the final PDF output. It's styled with a left border and different background.

### Mathematical Expressions

While basic formatting is supported, for complex math expressions, consider using LaTeX notation:

```latex
E = mc^2
```

## Advanced Features

### YAML Front Matter

This document uses YAML front matter to configure:

- Document title
- Theme selection
- Custom styling options

### Table of Contents

The `<!-- TOC -->` placeholder above automatically generates a table of contents based on your heading structure.

---

**Generated with Pdfy** - A modern Markdown to PDF converter
