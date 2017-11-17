package build

import (
	"os"

	"github.com/wsxiaoys/terminal/color"
)

var (
	// Verbose enables verbose logging
	Verbose bool
)

// LogFatal logs an error to stderr and exits with a non-zero exit code.
func LogFatal(err error) {
	color.Fprintf(os.Stderr, "@{r!}Fatal:@{r} %s\n", err)
	os.Exit(1)
}

// LogError logs an error to stderr.
func LogError(err error) {
	color.Fprintf(os.Stderr, "@{r!}Error:@{r} %s\n", err)
}

// LogWarning logs an error to stderr as a warning.
func LogWarning(err error) {
	color.Fprintf(os.Stderr, "@{y!}Warning:@{y} %s\n", err)
}

// LogWarningFor logs an error to stderr for a given label.
func LogWarningFor(err error, label string) {
	color.Fprintf(os.Stderr, "@{y!}Warning:@{y} %s@{|}{@.} [%s]\n", err, label)
}

// LogWarningForAt logs a warning to stderr for a given label and offset.
func LogWarningForAt(err error, label string, offset int) {
	color.Fprintf(os.Stderr, "@{y!}Warning:@{y} %s@{|}@{.} [%s]@@%d\n", err, label, offset)
}

// Trace logs an error to stderr.
func Trace(spec string, args ...interface{}) {
	color.Fprintf(os.Stderr, spec, args...)
}
