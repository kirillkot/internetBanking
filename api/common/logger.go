package common

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var (
	logger = NewLogger("common")
)

// LogFormatter ...
type LogFormatter struct {
	logrus.TextFormatter

	prefix string
}

// NewLogFormatter ...
func NewLogFormatter(prefix string) *LogFormatter {
	return &LogFormatter{
		prefix: prefix,
	}
}

// Format ...
func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	msg := entry.Message
	entry.Message = f.prefix + msg
	buf, err := f.TextFormatter.Format(entry)
	entry.Message = msg
	return buf, err
}

// NewLogger ...
func NewLogger(entry string) *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = NewLogFormatter(fmt.Sprintf("[%s]: ", entry))
	return logger
}
