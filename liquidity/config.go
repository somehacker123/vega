package liquidity

import (
	cfgencoding "code.vegaprotocol.io/vega/config/encoding"
	"code.vegaprotocol.io/vega/logging"
)

// namedLogger is the identifier for package and should ideally match the package name
// this is simply emitted as a hierarchical label e.g. 'api.grpc'.
const namedLogger = "liquidity"

// Config represents the configuration of the Liquidity service
type Config struct {
	Level cfgencoding.LogLevel `long:"log-level"`
}

// NewDefaultConfig creates an instance of the package specific configuration, given a
// pointer to a logger instance to be used for logging within the package.
func NewDefaultConfig() Config {
	return Config{
		Level: cfgencoding.LogLevel{Level: logging.InfoLevel},
	}
}