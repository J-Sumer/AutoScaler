package docker

import (
	"time"
	influxdb "github.com/J-Sumer/AutoScaler/velvet/influxDB"
)

func CollectAndAddMetrics() {
	cpu, memeory, ContainerMetrics := CollectMetric()
	// Export metrics 
	influxdb.AddMetricEntry(int(cpu), int(memeory), ContainerMetrics)
}

func MetricExporter(sec time.Duration) {

	for {
		time.Sleep(sec * time.Second)
		go CollectAndAddMetrics()
	}
}