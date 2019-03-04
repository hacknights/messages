package messages

import (
	"errors"
	"fmt"
	"strings"
)

type errorBuilder struct {
	sb strings.Builder
}

// NewErrorBuilder returns a pointer
// to an errorBuilder which wraps
// the strings.Builder
func NewErrorBuilder() *errorBuilder {
	return &errorBuilder{}
}

// Error returns the error with the
// built message
func (b *errorBuilder) Error() error {
	if 0 < b.sb.Len() {
		return errors.New(b.String())
	}
	return nil
}

// String returns the string value of
// the built error message removing
// the trailing return character
func (b *errorBuilder) String() string {
	return strings.TrimSuffix(b.sb.String(), "\n")
}

// WriteError unwraps the string value of
// the passed error and calls WriteString
func (b *errorBuilder) WriteError(err error) {
	b.WriteString(err.Error())
}

// WriteString adds a return character to
// the end of the passed string value and
// calls the WriteString method on the
// wrapped strings.Builder
func (b *errorBuilder) WriteString(err string) {
	b.sb.WriteString(fmt.Sprintf("%s\n", err))
}
