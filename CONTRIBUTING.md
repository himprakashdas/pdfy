# Contributing to Pdfy 🤝

Thank you for your interest in contributing to Pdfy! We welcome contributions from developers of all skill levels. This document provides guidelines and information to help you contribute effectively.

## 📋 Table of Contents

- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Code Style & Standards](#code-style--standards)
- [Contribution Workflow](#contribution-workflow)
- [Types of Contributions](#types-of-contributions)
- [Testing Guidelines](#testing-guidelines)
- [Documentation](#documentation)
- [Community Guidelines](#community-guidelines)

## 🚀 Getting Started

### Prerequisites

- **Go 1.21+** - [Install Go](https://golang.org/doc/install)
- **Git** - [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- **Chrome/Chromium** - Required for PDF generation
- **Make** (optional) - For using Makefile commands

### Quick Start

1. **Fork the repository** on GitHub
2. **Clone your fork**:
   ```bash
   git clone https://github.com/yourusername/pdfy.git
   cd pdfy
   ```
3. **Add upstream remote**:
   ```bash
   git remote add upstream https://github.com/originalowner/pdfy.git
   ```
4. **Install dependencies**:
   ```bash
   go mod download
   ```
5. **Build and test**:
   ```bash
   go build -o pdfy
   ./pdfy convert resources/source_markdown.MD
   ```

## 🛠️ Development Setup

### Project Structure

```
pdfy/
├── cmd/                    # CLI commands and subcommands
│   ├── root.go            # Root command configuration
│   ├── convert.go         # Single file conversion
│   ├── batch.go           # Batch processing
│   └── watch.go           # Watch mode
├── internal/converter/     # Core conversion logic
│   ├── config.go          # Configuration types
│   ├── converter.go       # Main conversion engine
│   ├── templates.go       # Template loading
│   ├── templates/         # HTML templates
│   │   └── default.html   # Default template
│   └── themes/            # CSS themes
│       ├── light.css      # Light theme
│       └── dark.css       # Dark theme
├── resources/             # Sample markdown files
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
├── main.go               # Application entry point
└── README.md             # Project documentation
```

### Environment Setup

```bash
# Set up development environment
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org

# Optional: Enable Go modules
export GO111MODULE=on

# For debugging Chrome issues
export CHROME_BIN=/path/to/chrome
```

### Build Commands

```bash
# Build for current platform
go build -o pdfy

# Build for all platforms
GOOS=linux GOARCH=amd64 go build -o pdfy-linux-amd64
GOOS=darwin GOARCH=amd64 go build -o pdfy-darwin-amd64
GOOS=windows GOARCH=amd64 go build -o pdfy-windows-amd64.exe

# Build with version info
go build -ldflags "-X main.version=v1.0.0" -o pdfy
```

## 📝 Code Style & Standards

### Go Style Guidelines

We follow standard Go conventions:

- **gofmt** - All code must be formatted with `gofmt`
- **go vet** - All code must pass `go vet`
- **golint** - Follow `golint` recommendations
- **Effective Go** - Follow [Effective Go](https://golang.org/doc/effective_go.html) principles

### Code Formatting

```bash
# Format all Go files
go fmt ./...

# Run linter
go vet ./...

# Optional: Use golangci-lint for comprehensive linting
golangci-lint run
```

### Naming Conventions

- **Packages**: Short, lowercase, single word when possible
- **Functions**: CamelCase, exported functions start with capital letter
- **Variables**: camelCase for local variables, CamelCase for exported
- **Constants**: SCREAMING_SNAKE_CASE or CamelCase for exported

### Error Handling

```go
// Good: Wrap errors with context
if err != nil {
    return fmt.Errorf("failed to convert markdown: %w", err)
}

// Good: Use specific error types when appropriate
type ConversionError struct {
    Message string
    Err     error
}

func (e *ConversionError) Error() string {
    return fmt.Sprintf("conversion error: %s", e.Message)
}
```

### Documentation

- All exported functions, types, and constants must have comments
- Comments should be complete sentences starting with the item name
- Use godoc-style comments

```go
// Convert converts markdown content to PDF format.
// It returns an error if the conversion fails.
func (c *Converter) Convert() error {
    // Implementation...
}
```

## 🔄 Contribution Workflow

### 1. Create an Issue

Before starting work, create or find an existing issue:

- **Bug reports**: Use the bug report template
- **Feature requests**: Use the feature request template
- **Questions**: Use GitHub Discussions

### 2. Create a Branch

```bash
# Update your fork
git checkout main
git pull upstream main

# Create a feature branch
git checkout -b feature/your-feature-name

# Or for bug fixes
git checkout -b fix/issue-description
```

### 3. Make Changes

- Keep commits small and focused
- Write clear commit messages
- Test your changes thoroughly

### 4. Commit Guidelines

Use conventional commit format:

```
type(scope): description

[optional body]

[optional footer]
```

Types:

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

Examples:

```bash
git commit -m "feat(converter): add dark theme support"
git commit -m "fix(cli): resolve batch processing crash on Windows"
git commit -m "docs(readme): update installation instructions"
```

### 5. Submit Pull Request

```bash
# Push your branch
git push origin feature/your-feature-name

# Create pull request on GitHub
```

### Pull Request Guidelines

- **Title**: Clear, descriptive title
- **Description**: Explain what changes you made and why
- **Testing**: Describe how you tested your changes
- **Screenshots**: Include screenshots for UI changes
- **Breaking changes**: Clearly document any breaking changes

## 📁 Types of Contributions

### 🐛 Bug Fixes

1. **Reproduce the bug** - Create a minimal test case
2. **Write a test** that demonstrates the bug
3. **Fix the bug** while ensuring the test passes
4. **Verify** the fix doesn't break existing functionality

### ✨ New Features

1. **Discuss the feature** in an issue first
2. **Design the API** - consider backward compatibility
3. **Implement** with tests and documentation
4. **Update** relevant documentation

### 📚 Documentation

- README improvements
- Code comments and godoc
- Usage examples
- Wiki pages
- Blog posts or tutorials

### 🎨 Themes and Templates

- New CSS themes in `internal/converter/themes/`
- HTML templates in `internal/converter/templates/`
- Follow existing naming conventions
- Include examples and documentation

## 🧪 Testing Guidelines

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run specific tests
go test -run TestConverterBasic ./internal/converter/
```

### Writing Tests

```go
func TestConverter_Convert(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
        wantErr  bool
    }{
        {
            name:     "basic markdown",
            input:    "# Hello\n\nWorld",
            expected: "<h1>Hello</h1>\n<p>World</p>",
            wantErr:  false,
        },
        // More test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            c := New(&Config{})
            result, err := c.markdownToHTML([]byte(tt.input))

            if tt.wantErr && err == nil {
                t.Error("expected error but got none")
            }

            if !tt.wantErr && err != nil {
                t.Errorf("unexpected error: %v", err)
            }

            if result != tt.expected {
                t.Errorf("expected %q, got %q", tt.expected, result)
            }
        })
    }
}
```

### Integration Tests

Create test markdown files in `testdata/` directory:

```bash
testdata/
├── basic.md
├── with-frontmatter.md
├── complex-document.md
└── expected/
    ├── basic.pdf
    ├── with-frontmatter.pdf
    └── complex-document.pdf
```

### Manual Testing

```bash
# Test different scenarios
./pdfy convert testdata/basic.md -o test-output/basic.pdf
./pdfy batch "testdata/*.md" --output-dir test-output/
./pdfy convert testdata/complex.md --theme dark
```

## 📖 Documentation

### Code Documentation

- Use clear, descriptive function and variable names
- Add comments for complex logic
- Document public APIs with examples

### User Documentation

- Update README.md for new features
- Add usage examples
- Update help text in CLI commands

### API Documentation

```go
// Example generates a complete API documentation example.
//
// Parameters:
//   - config: Configuration for the conversion process
//   - theme: Theme name (light, dark, or custom path)
//
// Returns the generated PDF content and any error encountered.
//
// Example:
//   config := &Config{InputPath: "doc.md", OutputPath: "doc.pdf"}
//   err := converter.Convert(config)
//   if err != nil {
//       log.Fatal(err)
//   }
func Example(config *Config, theme string) error {
    // Implementation...
}
```

## 🤝 Community Guidelines

### Code of Conduct

- Be respectful and inclusive
- Welcome newcomers and help them learn
- Focus on constructive feedback
- Respect different viewpoints and experiences

### Communication

- **GitHub Issues**: Bug reports, feature requests
- **GitHub Discussions**: Questions, ideas, general discussion
- **Pull Requests**: Code review and collaboration

### Getting Help

- Check existing issues and documentation first
- Provide minimal, reproducible examples
- Include relevant system information (OS, Go version, etc.)
- Be patient and respectful

## 🏆 Recognition

Contributors are recognized in several ways:

- **Contributors** section in README.md
- **Release notes** for significant contributions
- **GitHub discussions** highlighting valuable contributions

## 📋 Checklist for Contributors

Before submitting a pull request:

- [ ] Code follows Go style guidelines
- [ ] All tests pass (`go test ./...`)
- [ ] New features include tests
- [ ] Documentation is updated
- [ ] Commit messages follow conventional format
- [ ] No breaking changes (or clearly documented)
- [ ] Pull request description is clear and complete

## 📞 Questions?

- 💬 **GitHub Discussions**: For questions and general discussion
- 🐛 **GitHub Issues**: For bug reports and feature requests
- 📧 **Email**: maintainer@pdfy.dev (for security issues)

---

Thank you for contributing to Pdfy! 🙏 Your help makes this project better for everyone.
