package awsapicallviacloudtrail

type CpuOptions struct {
    ThreadsPerCore float64 `json:"threadsPerCore,omitempty"`
    CoreCount float64 `json:"coreCount,omitempty"`
}

func (c *CpuOptions) SetThreadsPerCore(threadsPerCore float64) {
    c.ThreadsPerCore = threadsPerCore
}

func (c *CpuOptions) SetCoreCount(coreCount float64) {
    c.CoreCount = coreCount
}
