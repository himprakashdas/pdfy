# Examples

This directory contains example Markdown files that demonstrate various features of Pdfy.

## Available Examples

### `example.md`

A comprehensive example showcasing all major Pdfy features:

- YAML front matter configuration
- Table of Contents generation
- Syntax highlighting
- Tables and lists
- Blockquotes and formatting
- Custom styling options

## Usage

Convert any example to PDF:

```bash
# Basic conversion
pdfy convert examples/example.md

# With custom output location
pdfy convert examples/example.md --output my-document.pdf

# With custom CSS
pdfy convert examples/example.md --css custom-styles.css
```

## Creating Your Own Examples

Feel free to add more examples to this directory. When creating examples:

1. Use descriptive filenames
2. Include YAML front matter when demonstrating configuration
3. Add comments to explain special features
4. Keep examples focused on specific use cases

## Testing

You can use these examples for testing Pdfy functionality during development without worrying about cleanup - the `.gitignore` file will automatically ignore any generated test PDFs.
