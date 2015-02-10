package core

type (
	Flow int
)

const (
	UnknownFlow Flow = iota
	SequentialFlow
	ParallelFlow
	DelayedFlow
)

var (
	flowSymbols = map[string]Flow{
		"s": SequentialFlow,
		"p": ParallelFlow,
		"d": DelayedFlow,
	}
	flowNames = map[Flow]string{
		UnknownFlow:    "Unknown",
		SequentialFlow: "Sequential",
		ParallelFlow:   "Parallel",
		DelayedFlow:    "Delayed",
	}
)

func FlowOf(f string) Flow {
	if flow, ok := flowSymbols[f]; ok {
		return flow
	} else {
		return UnknownFlow
	}
}

func (f Flow) String() string {
	name, _ := flowNames[f]
	return name
}
