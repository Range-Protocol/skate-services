package metrics

import (
	elMetrics "github.com/Layr-Labs/eigensdk-go/metrics"
)

type Metrics struct {
	elMetrics.Metrics
}

// NOTE: upgrade to custom implementations in future versions to support multichain features of Skate AVS
type EigenMetrics struct {
  elMetrics.EigenMetrics
}
