package zapdriver

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSourceLocation(t *testing.T) {
	t.Parallel()

	got := SourceLocation(runtime.Caller(0)).Interface.(*source)

	assert.Regexp(t, "/zapdriver/source_test.go", got.File)
	assert.Equal(t, "13", got.Line)
	assert.Regexp(t, "/zapdriver.TestSourceLocation$", got.Function)
}

func TestNewSource(t *testing.T) {
	t.Parallel()

	got := newSource(runtime.Caller(0))

	assert.Regexp(t, "/zapdriver/source_test.go", got.File)
	assert.Equal(t, "23", got.Line)
	assert.Regexp(t, "/zapdriver.TestNewSource$", got.Function)
}
