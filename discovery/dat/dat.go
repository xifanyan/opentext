package dat

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Reader is a custom CSV reader that supports configurable field separators and text qualifiers.
type Reader struct {
	scanner       *bufio.Scanner
	separator     rune
	qualifier     rune
	trimSpace     bool
	fieldsPerLine int
}

// Writer is a custom CSV writer that supports configurable field separators and text qualifiers.
type Writer struct {
	writer        *bufio.Writer
	separator     rune
	qualifier     rune
	alwaysQualify bool
}

// Option is a function type that configures a Reader or Writer.
type Option func(interface{}) error

// NewReader creates a new custom CSV reader with optional configurations.
func NewReader(r io.Reader, options ...Option) *Reader {
	reader := &Reader{
		scanner:       bufio.NewScanner(r),
		separator:     '\u0014', // Default separator
		qualifier:     '\u00fe', // Default qualifier
		trimSpace:     false,
		fieldsPerLine: -1, // Default: variable number of fields
	}

	for _, option := range options {
		if err := option(reader); err != nil {
			return nil
		}
	}

	return reader
}

// NewWriter creates a new custom CSV writer with optional configurations.
func NewWriter(w io.Writer, options ...Option) (*Writer, error) {
	writer := &Writer{
		writer:    bufio.NewWriter(w),
		separator: ',', // Default separator
		qualifier: '"', // Default qualifier
	}

	for _, option := range options {
		if err := option(writer); err != nil {
			return nil, err
		}
	}

	return writer, nil
}

// WithSeparator sets the field separator.
func WithSeparator(separator rune) Option {
	return func(i interface{}) error {
		switch v := i.(type) {
		case *Reader:
			v.separator = separator
		case *Writer:
			v.separator = separator
		default:
			return fmt.Errorf("unknown type: %T", i)
		}
		return nil
	}
}

// WithQualifier sets the text qualifier.
func WithQualifier(qualifier rune) Option {
	return func(i interface{}) error {
		switch v := i.(type) {
		case *Reader:
			v.qualifier = qualifier
		case *Writer:
			v.qualifier = qualifier
		default:
			return fmt.Errorf("unknown type: %T", i)
		}
		return nil
	}
}

// WithTrimSpace enables or disables trimming leading and trailing spaces.
func WithTrimSpace(trimSpace bool) Option {
	return func(i interface{}) error {
		if reader, ok := i.(*Reader); ok {
			reader.trimSpace = trimSpace
		}
		return nil
	}
}

// WithFieldsPerLine sets the expected number of fields per line.
func WithFieldsPerLine(fieldsPerLine int) Option {
	return func(i interface{}) error {
		if reader, ok := i.(*Reader); ok {
			reader.fieldsPerLine = fieldsPerLine
		}
		return nil
	}
}

func WithAlwaysQualify(alwaysQualify bool) Option {
	return func(i interface{}) error {
		if writer, ok := i.(*Writer); ok {
			writer.alwaysQualify = alwaysQualify
		}
		return nil
	}
}

func (r *Reader) NextLine() error {
	if !r.scanner.Scan() {
		return r.scanner.Err()
	}
	return nil
}

// Read reads a single record from the CSV.
func (r *Reader) Read() ([]string, error) {

	// handle the case where r.scanner.Scan() returns false
	if !r.scanner.Scan() {
		// Check if there was an error, If so, return nil and the error
		err := r.scanner.Err()
		if err != nil {
			return nil, err
		}
		// If there was no error, return nil and io.EOF
		return nil, io.EOF
	}
	line := r.scanner.Text()

	fields := r.ParseLine(line)
	if r.fieldsPerLine > 0 && len(fields) != r.fieldsPerLine {
		return nil, fmt.Errorf("wrong number of fields in line: expected %d, got %d", r.fieldsPerLine, len(fields))
	}

	return fields, nil
}

func (r *Reader) ParseLine(line string) []string {
	var fields []string
	var inField bool
	var field strings.Builder

	for _, char := range line {
		if char == r.qualifier {
			inField = !inField
		} else if char == r.separator && !inField {
			fields = append(fields, r.trim(field.String()))
			field.Reset()
		} else {
			field.WriteRune(char)
		}
	}

	// Append the last field
	if field.Len() > 0 {
		fields = append(fields, r.trim(field.String()))
	}

	return fields
}

func (r *Reader) trim(s string) string {
	if r.trimSpace {
		return strings.TrimSpace(s)
	}
	return s
}

// Write writes a single record to the CSV.
func (w *Writer) Write(record []string) error {
	for i, field := range record {
		if i > 0 {
			w.writer.WriteRune(w.separator)
		}
		w.writeField(field)
	}
	w.writer.WriteRune('\n')
	return w.writer.Flush()
}

func (w *Writer) writeField(field string) {
	needsQualifier := w.alwaysQualify || strings.ContainsRune(field, w.separator) || strings.ContainsRune(field, w.qualifier)

	if needsQualifier {
		w.writer.WriteRune(w.qualifier)
		for _, char := range field {
			if char == w.qualifier {
				w.writer.WriteRune(w.qualifier)
			}
			w.writer.WriteRune(char)
		}
		w.writer.WriteRune(w.qualifier)
	} else {
		w.writer.WriteString(field)
	}
}
