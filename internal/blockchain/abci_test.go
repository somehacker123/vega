package blockchain

import (
	"testing"
	"vega/internal/vegatime"
	execution "vega/internal/execution/mocks"
	"github.com/stretchr/testify/assert"
	"vega/internal/logging"
)

func TestNewAbciApplication(t *testing.T) {
	ex := &execution.Engine{}
	vt := vegatime.NewTimeService(nil)

	logger := logging.NewLogger()
	logger.InitConsoleLogger(logging.DebugLevel)
	logger.AddExitHandler()

	config := NewConfig(logger)
	stats := NewStats()
	chain := NewAbciApplication(config, ex, vt, stats)
	assert.Equal(t, uint64(0), chain.height)
}
