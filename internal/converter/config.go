package converter

import (
	"fmt"
	"time"
)

// Config holds the configuration for the conversion process
type Config struct {
	InputPath    string
	OutputPath   string
	TemplateName string
	CSSPath      string
	Theme        string
}

// FrontMatter represents YAML front matter configuration
type FrontMatter struct {
	Title    string `yaml:"title"`
	Theme    string `yaml:"theme"`
	Template string `yaml:"template"`
	CSS      string `yaml:"css"`
}

// ConversionError represents an error during conversion
type ConversionError struct {
	LineNumber int
	Message    string
	Snippet    string
	Err        error
}

func (e *ConversionError) Error() string {
	if e.LineNumber > 0 {
		return fmt.Sprintf("line %d: %s - %s", e.LineNumber, e.Message, e.Snippet)
	}
	return e.Message
}

// ConversionStats holds statistics about the conversion process
type ConversionStats struct {
	StartTime    time.Time
	EndTime      time.Time
	InputSize    int64
	OutputSize   int64
	PageCount    int
	ProcessingMS int64
}
