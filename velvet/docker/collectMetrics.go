package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	stypes "github.com/J-Sumer/AutoScaler/velvet/docker/types"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func getStats(cli *client.Client, ctx context.Context,containerID string, channel chan stypes.Metrics) {
	stats, err := cli.ContainerStats(ctx, containerID, false)
	if err != nil {
		panic(err)
	}
	bytes, err := io.ReadAll(stats.Body)
	if err != nil {
		panic(err)
	}
	var result stypes.CompleteStats
	json.Unmarshal(bytes, &result)

	// Calculate Memory usage
	used_memory := result.MemoryStats.Usage - result.MemoryStats.Stats.Cache
	available_memory := result.MemoryStats.Limit
	MemoryUsage := (used_memory/available_memory) * 100

	// Calculate CPU usage
	cpu_delta := result.CpuStats.CpuUsage.TotalUsage - result.PreCpuStats.CpuUsage.TotalUsage
	system_cpu_delta := result.CpuStats.SystemCPUUsage - result.PreCpuStats.SystemCPUUsage
	num_cpus := result.CpuStats.OnlineCPUs
	CPUUsage := (cpu_delta/system_cpu_delta) * num_cpus * 100.0

	// fmt.Println("Memory usage", MemoryUsage)
	// fmt.Println("CPU Usage", CPUUsage)
	metric := stypes.Metrics{
		ContainerId: containerID,
		CPU: CPUUsage,
		Memory: MemoryUsage,
	}
	channel <- metric
}

func CollectMetric() (float32, float32, stypes.ContainerStats) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	ctx := context.Background()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(len(containers))

	channel := make(chan stypes.Metrics)

	for _, container := range containers {
		containerID := container.ID[:10]
		// fmt.Printf("%s %s\n", containerID, container.Image)
		go getStats(cli , ctx, containerID, channel)
	}

	var ContainerMetrics stypes.ContainerStats
	var MemoryMetric float32
	var CPUMetric float32
	for i := 0; i < len(containers); i++ {
		// fmt.Println(i)
		metric := <- channel
		ContainerMetrics.AllMetrics = append(ContainerMetrics.AllMetrics, metric)
		MemoryMetric += metric.Memory
		CPUMetric += metric.CPU
	}
	// fmt.Println("Memory", (MemoryMetric/float32(len(containers))))
	// fmt.Println("CPU", (CPUMetric/float32(len(containers))))

	// send data to influxDB
	return CPUMetric, MemoryMetric, ContainerMetrics
}
