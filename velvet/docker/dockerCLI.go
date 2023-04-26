package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	stypes "github.com/J-Sumer/AutoScaler/velvet/docker/types"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	spec "github.com/opencontainers/image-spec/specs-go/v1"
)

func CreateContainer() string {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	ctx := context.Background()
	if err != nil {
		panic(err)
	}

	containerConfig := container.Config{
		Image: "jsumermaduru/rubis",
	}

	containerResources := container.Resources{
		CPUQuota: 20000,
	}

	hostConfig := container.HostConfig{
		Resources: containerResources,
		AutoRemove: true,
	}

	container, err := cli.ContainerCreate(ctx, &containerConfig, &hostConfig, &network.NetworkingConfig{}, &spec.Platform{}, "")
	if err != nil {
		panic(err)
	}
	fmt.Print(container.ID)

	return container.ID
}

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

	// fmt.Println(len(containers))

	channel := make(chan stypes.Metrics)

	for _, container := range containers {
		containerID := container.ID[:12]
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
		MemoryMetric += (metric.Memory)
		CPUMetric += (metric.CPU * 5)
	}
	// fmt.Println("Memory", (MemoryMetric/float32(len(containers))))
	// fmt.Println("CPU", (CPUMetric/float32(len(containers))))

	// send data to influxDB
	return (CPUMetric/float32(len(containers))), (MemoryMetric/float32(len(containers))) , ContainerMetrics
}

func RunningContainers() []string {

	var containerList []string
	cli, err := client.NewClientWithOpts(client.FromEnv)
	ctx := context.Background()
	if err != nil {
		panic(err)
	}

	fmt.Println("Getting list of running containers")

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		containerList = append(containerList, container.ID[:12])
	}

	return containerList
}