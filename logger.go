package slf4go //github.com/less4go/slf4go

import (
	conf "github.com/less4go/slf4go/config"
	"github.com/less4go/slf4go/provider"
	"github.com/less4go/slf4go/types"
)

// Logger var for logger provider
var Logger types.Logger

// Logger interface

func init() {
	Logger = initLogger()
}

func initLogger() types.Logger {

	logConfig := &conf.LogConfig{}
	logConfig.ReadConfig()

	var prov types.Bootstrap
	switch logConfig.Log.Provider {
	case "zap":
		prov = &provider.ZapLoggerProvider{}
		break
	case "logrus":
		prov = &provider.LogrusLoggerProvider{}
		break
	default:
		prov = &provider.ZapLoggerProvider{}
	}
	//

	return prov.BootLogger(logConfig)
}
