package docker

import (
	"time"
	"github.com/J-Sumer/AutoScaler/velvet/routes"
)

func CollectAndAddMetrics() {
	cpu, memeory := CollectMetric()
	// Export metrics 
	routes.AddMetricEntry(int(cpu), int(memeory))
}

func MetricExporter(sec time.Duration) {

	for {
		time.Sleep(sec * time.Second)
		go CollectAndAddMetrics()
	}
}