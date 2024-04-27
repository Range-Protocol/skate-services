package monitor

import (
	"context"
	"time"

	"skate.org/avs/lib/logging"
	"skate.org/avs/lib/network"
	"skate.org/avs/lib/on-chain/backend"
)

// Monitor interface defines common monitor operations
type Monitor interface {
	Start()
	MonitorOnce(ctx context.Context, network network.Network, backends map[uint64]backend.Backend) error
	MonitorForever(ctx context.Context, network network.Network, backends map[uint64]backend.Backend, processName string, interval time.Duration)
}

// NOTE: specific monitor to be implemented by each process
type MonitorBase struct{}

var _ Monitor = (*MonitorBase)(nil)

func (m *MonitorBase) Start() {}

func (m *MonitorBase) MonitorOnce(ctx context.Context, network network.Network, backends map[uint64]backend.Backend) error {
	return nil
}

// MonitorForever implements Monitor interface
func (m *MonitorBase) MonitorForever(
	ctx context.Context, network network.Network,
	backends map[uint64]backend.Backend, processName string, interval time.Duration,
) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	logger := logging.NewLoggerWithConsoleWriter()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ctx, cancel := context.WithTimeout(ctx, interval)
			defer cancel()
			err := m.MonitorOnce(ctx, network, backends)
			if err != nil {
				logger.Error("Monitor failed, retrying next interval ...", "name", processName, "error", err)
			}
		}
	}
}

