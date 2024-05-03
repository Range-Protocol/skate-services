package publish

import (
	// avsMemcache "skatechain.org/relayer/db/avs/mem"
	// skateappMemcache "skatechain.org/relayer/db/skateapp/mem"
	"skatechain.org/lib/logging"
)

// NOTE: use shared in-mem cache if facing performance is a bottleneck.
// To be consider in future versions
var (
	Verbose       = true
	// taskCache     = skateappMemcache.NewCache(100 * 1024 * 1024) // 100MB
	// operatorCache = avsMemcache.NewCache(2 * 1024 * 1024)        // 2MB
	relayerLogger = logging.NewLoggerWithConsoleWriter()
)

