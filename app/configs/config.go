package configs

import (
	"flag"
	"strings"
	"time"

	"github.com/FZambia/viper-lite"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

const (
	ConfigDebug       = "debug"
	ConfigLogLevel    = "log-level"
	ConfigServiceName = "lrn"

	ConfigMongoDSN = "mongo-dsn"
	ConfigMongoDB  = "mongo-db"

	ConfigServeAddr         = "serve-addr"
	ConfigInternalServeAddr = "internal-serve-addr"
	ConfigRequestTimeout    = "request-timeout"

	ConfigLeaseMaxLimit = "lease-max-limit"
)

func init() {
	// base config
	pflag.Bool(ConfigDebug, true, "enable debug mode")
	pflag.Int(ConfigLogLevel, int(zap.DebugLevel), "")
	pflag.String(ConfigServiceName, "proxy-manager-sv", "")

	// common
	pflag.Int(ConfigLeaseMaxLimit, 10, "")

	// server
	pflag.String(ConfigServeAddr, ":80", "")
	pflag.String(ConfigInternalServeAddr, ":8081", "")
	pflag.Duration(ConfigRequestTimeout, 10*time.Second, "")

	// mongodb
	pflag.String(ConfigMongoDSN, "mongodb://127.0.0.1:27017", "")
	pflag.String(ConfigMongoDB, "db-name", "")

}

func InitConfig() {
	pflag.Parse()

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	_ = viper.BindPFlags(pflag.CommandLine)
	_ = flag.CommandLine.Parse([]string{})
}
