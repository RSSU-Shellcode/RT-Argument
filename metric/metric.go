package metric

import (
	"github.com/RTS-Framework/GRT-Develop/types"
)

// Metrics contains status about runtime submodules.
type Metrics struct {
	Library  LTStatus
	Memory   MTStatus
	Thread   TTStatus
	Resource RTStatus
	Argument ASStatus
	Storage  ISStatus
	Detector DTStatus
	Watchdog WDStatus
	Sysmon   SMStatus
	Shield   SDStatus
	Core     RTCore
	Proc     RTProc
	Sleep    RTSleep
}

// LTStatus contains status about library tracker.
type LTStatus struct {
	NumModules   int64
	NumLoadCalls int64
	NumFreeCalls int64
}

// MTStatus contains status about memory tracker.
type MTStatus struct {
	NumGlobals int64
	NumLocals  int64
	NumBlocks  int64
	NumRegions int64
	NumPages   int64
	NumHeaps   int64
	NumRWXs    int64
	TotalAlloc int64
	PeakAlloc  int64
}

// TTStatus contains status about thread tracker.
type TTStatus struct {
	NumThreads   int64
	NumTLSIndex  int64
	NumCreated   int64
	NumExited    int64
	NumLocked    int64
	NumSuspended int64
}

// RTStatus contains status about resource tracker.
type RTStatus struct {
	NumMutexs         int64
	NumEvents         int64
	NumSemaphores     int64
	NumWaitableTimers int64
	NumFiles          int64
	NumDirectories    int64
	NumIOCPs          int64
	NumRegKeys        int64
	NumSockets        int64
}

// ASStatus contains status about argument store.
type ASStatus struct {
	NumItems  int32
	NumErased int32
	TotalSize int64
}

// ISStatus contains status about in-memory storage.
type ISStatus struct {
	NumItems  int64
	TotalSize int64
}

// DTStatus contains status about detector.
type DTStatus struct {
	IsEnabled        types.BOOL
	HasDebugger      types.BOOL
	HasMemoryScanner types.BOOL
	InSandbox        types.BOOL
	InEmulator       types.BOOL
	InVirtualMachine types.BOOL
	IsAccelerated    types.BOOL
	SafeRank         int32
	NumDetectCalls   int64
}

// WDStatus contains status about watchdog.
type WDStatus struct {
	IsEnabled types.BOOL
	Reserved  int32
	NumKick   int64
	NumNormal int64
	NumReset  int64
}

// SMStatus contains status about sysmon.
type SMStatus struct {
	IsEnabled  types.BOOL
	Reserved   int32
	NumNormal  int64
	NumRecover int64
	NumPanic   int64
}

// SDStatus contains status about shield.
type SDStatus struct {
	EntryPoint  uintptr
	BaseAddress uintptr
	Source      int64
}

// RTCore contains metric about runtime core.
type RTCore struct {
	Uptime       int64
	InitElapsed  int64
	SecurityMode types.BOOL
	IsHealthy    types.BOOL
}

// RTProc contains metric about runtime GetProcAddress.
type RTProc struct {
	NumCalls    int64
	NumRedirect int64
	NumFallback int64
	NumRTMethod int64
	NumRawProc  int64
}

// RTSleep contains metric about runtime sleep.
type RTSleep struct {
	NumCalls         int64
	LastError        int32
	Reserved         int32
	LastPreElapsed   int32
	LastPostElapsed  int32
	TotalPreElapsed  int64
	TotalPostElapsed int64
	MinPreElapsed    int32
	MinPostElapsed   int32
	MaxPreElapsed    int32
	MaxPostElapsed   int32
	AvgPreElapsed    int32
	AvgPostElapsed   int32
}
