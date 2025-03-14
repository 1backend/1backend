package registryservice_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	registryservice "github.com/openorch/openorch/server/internal/services/registry"
)

// output for
// nvidia-smi --query-gpu=name,temperature.gpu,utilization.gpu,memory.total,memory.used,power.draw,power.limit,driver_version,pci.bus_id,compute_mode,pstate --format=csv,noheader,nounits
const nvidiaSmiOutput = `
NVIDIA GeForce RTX 3090, 48, 0, 24576, 2592, 22.88, 350.00, 535.183.01, 00000000:04:00.0, Default, P8
NVIDIA GeForce RTX 3090, 52, 0, 24576, 2600, 26.85, 350.00, 535.183.01, 00000000:2B:00.0, Default, P8
`

func TestNvidiaSmiOutput(t *testing.T) {
	ns := registryservice.RegistryService{URL: "testhost:58231"}

	gpus, err := ns.ParseNvidiaSmiOutput(nvidiaSmiOutput)
	require.NoError(t, err)

	require.Equal(t, 2, len(gpus))

	gpu1 := gpus[0]
	require.Equal(t, "testhost:58231:0", gpu1.Id)
	require.Equal(t, 0, gpu1.IntraNodeId)
	require.Equal(t, "NVIDIA GeForce RTX 3090", gpu1.Name)
	require.Equal(t, "00000000:04:00.0", gpu1.BusId)
	require.Equal(t, 48.0, gpu1.Temperature)
	require.Equal(t, "P8", gpu1.PerformanceState)
	require.Equal(t, 22.88, gpu1.PowerUsage)
	require.Equal(t, 350.00, gpu1.PowerCap)
	require.Equal(t, 2592, gpu1.MemoryUsage)
	require.Equal(t, 24576, gpu1.MemoryTotal)
	require.Equal(t, float64(0), gpu1.GPUUtilization)
	require.Equal(t, "Default", gpu1.ComputeMode)

	gpu2 := gpus[1]
	require.Equal(t, "testhost:58231:1", gpu2.Id)
	require.Equal(t, 1, gpu2.IntraNodeId)
	require.Equal(t, "NVIDIA GeForce RTX 3090", gpu2.Name)
	require.Equal(t, "00000000:2B:00.0", gpu2.BusId)
	require.Equal(t, 52.0, gpu2.Temperature)
	require.Equal(t, "P8", gpu2.PerformanceState)
	require.Equal(t, 26.85, gpu2.PowerUsage)
	require.Equal(t, 350.00, gpu2.PowerCap)
	require.Equal(t, 2600, gpu2.MemoryUsage)
	require.Equal(t, 24576, gpu2.MemoryTotal)
	require.Equal(t, float64(0), gpu2.GPUUtilization)
	require.Equal(t, "Default", gpu2.ComputeMode)
}

const nvidiaSmiQueryOutput = `
some lines
CUDA Version                              : 12.2
some more lines
`

func TestNvidiaSmiQueryOutput(t *testing.T) {
	ns := registryservice.RegistryService{URL: "testhost:58231"}

	cudaVersion, err := ns.ParseNvidiaSmiQueryOutput(nvidiaSmiQueryOutput)
	require.NoError(t, err)
	require.Equal(t, "12.2", cudaVersion)
}
