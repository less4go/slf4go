package provider

import (
	"encoding/json"

	"github.com/less4go/slf4go/config"
	"github.com/less4go/slf4go/types"
	"go.uber.org/zap"
)

// ZapLoggerProvider struct
type ZapLoggerProvider struct {
	logger *zap.Logger
}

// BootLogger func for ZapLoggerProvider
func (*ZapLoggerProvider) BootLogger(logConfig *config.LogConfig) types.Logger {

	bytes, err := json.Marshal(logConfig.Log.Zap)
	if err != nil {
		panic(err)
	}

	cfg := zap.Config{}
	err = json.Unmarshal(bytes, &cfg)
	if err != nil {
		panic(err)
	}

	logger, err := cfg.Build(zap.AddCallerSkip(logConfig.Log.Zap.CallerSkip))
	defer logger.Sync()
	if err != nil {
		panic(err)
	}

	return &ZapLoggerProvider{logger}
}

// Debug fun for zap provider
func (provider *ZapLoggerProvider) Debug(msg string, fields ...types.Field) {
	provider.logger.Debug(msg, provider.buildFields(fields)...)
}

// Info fun for zap provider
func (provider *ZapLoggerProvider) Info(msg string, fields ...types.Field) {
	provider.logger.Info(msg, provider.buildFields(fields)...)
}

// Warn fun for zap provider
func (provider *ZapLoggerProvider) Warn(msg string, fields ...types.Field) {
	provider.logger.Warn(msg, provider.buildFields(fields)...)
}

// Error fun for zap provider
func (provider *ZapLoggerProvider) Error(msg string, fields ...types.Field) {
	provider.logger.Error(msg, provider.buildFields(fields)...)
}

// Fatal fun for zap provider
func (provider *ZapLoggerProvider) Fatal(msg string, fields ...types.Field) {
	provider.logger.Fatal(msg, provider.buildFields(fields)...)
}

// Panic fun for zap provider
func (provider *ZapLoggerProvider) Panic(msg string, fields ...types.Field) {

	provider.logger.Panic(msg, provider.buildFields(fields)...)
}

// With fun for zap provider
func (provider *ZapLoggerProvider) With(fields ...types.Field) types.Logger {
	copy := *provider
	copy.logger = provider.logger.With(provider.buildFields(fields)...)
	return &copy
}

func (*ZapLoggerProvider) buildFields(fields []types.Field) []zap.Field {
	var withFields []zap.Field

	for _, field := range fields {
		withFields = append(withFields, zap.Any(field.Key, field.Val))
	}

	return withFields
}
