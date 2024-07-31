package randgen

import (
	"errors"
	"os"
	"runtime"

	"github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/v4/mem"
)

const (
	envar_RANDGEN_DATA_MEMORY_LIMIT = "RANDGEN_API_DATA_LIMIT"
	divFactorPerCPU                 = 4
)

var (
	dataLimit = func() int {
		limitStr := os.Getenv(envar_RANDGEN_DATA_MEMORY_LIMIT)
		limit, err := humanize.ParseBytes(limitStr)
		if err != nil {
			divFactor := runtime.NumCPU() * divFactorPerCPU
			virtMem, err := mem.VirtualMemory()
			if err == nil {
				return int(virtMem.Available) / divFactor
			}
		}
		return int(limit)
	}()

	errDataLimitExceeded = errors.New("requested size is greater than configured/available memory")
)
