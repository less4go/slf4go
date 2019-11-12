package provider

import (
	"fmt"
	"io"
	"os"

	"github.com/less4go/slf4go/config"
	"github.com/less4go/slf4go/types"
	"github.com/sirupsen/logrus"
)

// LogrusLoggerProvider struct
type LogrusLoggerProvider struct {
	logger *logrus.Entry
}

// BootLogger func for LogrusLoggerProvider
func (provider *LogrusLoggerProvider) BootLogger(logConfig *config.LogConfig) types.Logger {
	logger := logrus.New()

	level, err := logrus.ParseLevel(logConfig.Log.Logrus.Level)
	if err != nil {
		panic(err)
	}
	logger.SetLevel(level)

	if logConfig.Log.Logrus.NoLock {
		logger.SetNoLock()
	}

	if logConfig.Log.Logrus.Formatter == "json" {
		logger.SetFormatter(logConfig.Log.Logrus.JSONFormatter)
	} else if logConfig.Log.Logrus.Formatter == "text" {
		logger.SetFormatter(logConfig.Log.Logrus.TextFormatter)
	}

	logger.SetOutput(provider.getOutout(logConfig.Log.Logrus.Output))

	logger.SetReportCaller(logConfig.Log.Logrus.ReportCaller)

	return &LogrusLoggerProvider{logrus.NewEntry(logger)}
}

// Debug fun for logrus provider
func (provider *LogrusLoggerProvider) Debug(msg string, fields ...*types.Field) {
	provider.logger.WithFields(provider.buildFields(fields)).Debug(msg)
}

// Info fun for logrus provider
func (provider *LogrusLoggerProvider) Info(msg string, fields ...*types.Field) {
	provider.logger.WithFields(provider.buildFields(fields)).Info(msg)
}

// Warn fun for logrus provider
func (provider *LogrusLoggerProvider) Warn(msg string, fields ...*types.Field) {
	provider.logger.WithFields(provider.buildFields(fields)).Warn(msg)
}

// Error fun for logrus provider
func (provider *LogrusLoggerProvider) Error(msg string, fields ...*types.Field) {
	provider.logger.WithFields(provider.buildFields(fields)).Error(msg)
}

// Fatal fun for logrus provider
func (provider *LogrusLoggerProvider) Fatal(msg string, fields ...*types.Field) {
	provider.logger.WithFields(provider.buildFields(fields)).Fatal(msg)
}

// Panic fun for logrus provider
func (provider *LogrusLoggerProvider) Panic(msg string, fields ...*types.Field) {
	provider.logger.WithFields(provider.buildFields(fields)).Panic(msg)
}

// With fun for logrus provider
func (provider *LogrusLoggerProvider) With(fields ...*types.Field) types.Logger {
	provider.logger = provider.logger.WithFields(provider.buildFields(fields))
	return provider
}

func (*LogrusLoggerProvider) buildFields(fields []*types.Field) logrus.Fields {
	withFields := make(logrus.Fields)
	for _, field := range fields {
		withFields[field.Key] = field.Val
	}
	return withFields
}

func (*LogrusLoggerProvider) getOutout(output string) io.Writer {

	var writer io.Writer = os.Stdout
	if len(output) > 0 {
		file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			writer = file
		} else {
			fmt.Println("Failed to log to file")
		}
	}
	return writer
}
