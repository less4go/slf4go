package config //github.com/less4go/slf4go/config

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var once sync.Once

type (
	// LogConfig struct
	LogConfig struct {
		Log struct {
			Provider string
			Zap      ZapConfig
			Logrus   LogrusConfig
		}
	}

	// ZapConfig struct
	ZapConfig struct {
		Level             string
		Development       bool
		DisableCaller     bool
		DisableStacktrace bool
		Sampling          struct {
			Initial    int
			Thereafter int
		}
		Encoding      string // "json" and "console"
		EncoderConfig struct {
			MessageKey      string
			LevelKey        string
			TimeKey         string
			NameKey         string
			CallerKey       string
			StacktraceKey   string
			LineEnding      string // \n
			LevelEncoder    string // capital, capitalColor, color
			TimeEncoder     string // rfc3339nano, rfc3339, iso8601, millis, nanos
			DurationEncoder string // string, nanos
			CallerEncoder   string // full
			NameEncoder     string // full
		}
		OutputPaths      []string
		ErrorOutputPaths []string
		InitialFields    map[string]interface{}
		CallerSkip       int
	}

	// LogrusConfig struct
	LogrusConfig struct {
		Level         string
		Output        string
		NoLock        bool
		Formatter     string
		ReportCaller  bool
		JSONFormatter *logrus.JSONFormatter
		TextFormatter *logrus.TextFormatter
	}
)

// ReadConfig func
func (config *LogConfig) ReadConfig() {
	once.Do(func() {

		viper := viper.New()
		viper.AutomaticEnv()

		logFile := viper.GetString("SLF4GO_FNAME")
		if len(logFile) <= 0 {
			logFile = "slf4go"
		}
		viper.SetConfigName(logFile)
		viper.AddConfigPath(".")

		// Zap default value setting.
		viper.SetDefault("log.zap.level", "info")
		viper.SetDefault("log.zap.encoding", "console")
		viper.SetDefault("log.zap.disableStacktrace", true)
		viper.SetDefault("log.zap.sampling.initial", 1)
		viper.SetDefault("log.zap.sampling.thereafter", 1)
		viper.SetDefault("log.zap.callerSkip", 1)
		viper.SetDefault("log.zap.outputPaths", "stdout")

		viper.SetDefault("log.zap.encoderConfig.messageKey", "msg")
		viper.SetDefault("log.zap.encoderConfig.levelKey", "level")
		viper.SetDefault("log.zap.encoderConfig.levelEncoder", "capitalColor")
		viper.SetDefault("log.zap.encoderConfig.lineEnding", "\n")
		viper.SetDefault("log.zap.encoderConfig.timeKey", "ts")
		viper.SetDefault("log.zap.encoderConfig.nameKey", "logger")
		viper.SetDefault("log.zap.encoderConfig.callerKey", "caller")
		viper.SetDefault("log.zap.encoderConfig.stacktraceKey", "stacktrace")
		viper.SetDefault("log.zap.encoderConfig.timeEncoder", "iso8601")

		// Logrus default value setting.
		viper.SetDefault("log.logrus.level", "info")
		viper.SetDefault("log.logrus.jsonFormatter.disableTimestamp", false)
		viper.SetDefault("log.logrus.textFormatter.disableTimestamp", false)

		viper.ReadInConfig()
		viper.Unmarshal(config)

	})
}
