package types

type Metrics struct {
	CPU float32
	Memory float32
}

// Memory stats
type Stats struct {
	Cache float32 `json:"cache"`
}

type MemoryStats struct {
	Stats Stats `json:"stats"`
	MaxUsage float32 `json:"max_usage"`
	Usage float32 `json:"usage"`
	Limit float32 `json:"limit"`
}

// CPU Stats
type CPUUsage struct {
	TotalUsage float32 `json:"total_usage"`
}

type CPUStats struct {
	CpuUsage CPUUsage `json:"cpu_usage"`
	SystemCPUUsage float32 `json:"system_cpu_usage"`
	OnlineCPUs float32 `json:"online_cpus"`
}

type CompleteStats struct {
    MemoryStats MemoryStats `json:"memory_stats"`
	CpuStats CPUStats `json:"cpu_stats"`
	PreCpuStats CPUStats `json:"precpu_stats"`
}