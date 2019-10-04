package zapdriver

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorReport(t *testing.T) {
	t.Parallel()

	got := ErrorReport(runtime.Caller(0)).Interface.(*reportContext)

	assert.Regexp(t, "/zapdriver/report_test.go", got.ReportLocation.File)
	assert.Equal(t, "13", got.ReportLocation.Line)
	assert.Regexp(t, "/zapdriver.TestErrorReport$", got.ReportLocation.Function)
}

func TestNewReportContext(t *testing.T) {
	t.Parallel()

	got := newReportContext(runtime.Caller(0))

	assert.Regexp(t, "/zapdriver/report_test.go", got.ReportLocation.File)
	assert.Equal(t, "23", got.ReportLocation.Line)
	assert.Regexp(t, "/zapdriver.TestNewReportContext$", got.ReportLocation.Function)
}
