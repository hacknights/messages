package messages

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hacknights/testing/assert"
)

func TestString_0(t *testing.T) {
	b := NewErrorBuilder()
	if s := b.String(); s != "" {
		assert.Fail(t, "", s)
	}
}

const expectedHello = "hello"

func TestString_1(t *testing.T) {
	b := NewErrorBuilder()
	b.WriteString(expectedHello)
	if s := b.String(); s != expectedHello {
		assert.Fail(t, expectedHello, s)
	}
}

const expectedWorld = "world"

var expectedHelloWorld = fmt.Sprintf("%s\n%s", expectedHello, expectedWorld)

func TestString_2(t *testing.T) {
	b := NewErrorBuilder()
	b.WriteString(expectedHello)
	b.WriteString(expectedWorld)
	if s := b.String(); s != expectedHelloWorld {
		assert.Fail(t, expectedHelloWorld, s)
	}
}

func TestError_0(t *testing.T) {
	b := NewErrorBuilder()
	if err := b.Error(); err != nil {
		assert.Error(t, nil, err)
	}
}

func TestError_1(t *testing.T) {
	b := NewErrorBuilder()
	b.WriteError(errors.New(expectedHello))
	if err := b.Error(); err.Error() != expectedHello {
		assert.Error(t, expectedHello, err)
	}
}

func TestError_2(t *testing.T) {
	b := NewErrorBuilder()
	b.WriteError(errors.New(expectedHello))
	b.WriteError(errors.New(expectedWorld))
	if err := b.Error(); err.Error() != expectedHelloWorld {
		assert.Error(t, expectedHelloWorld, err)
	}
}
