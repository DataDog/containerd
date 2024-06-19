package plugin

import (
	"fmt"

	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type NameSampler struct {
	// allow is a set of names that should be sampled.
	// Uses a map of empty structs for O(1) lookups and no memory overhead.
	allow map[string]struct{}
}

// NameBased returns a Sampler that samples every span having a certain name.
// It should be used in conjunction with the ParentBased sampler so that the child spans are also sampled.
func NameBased(allowedNames []string) NameSampler {
	allowedNamesMap := make(map[string]struct{}, len(allowedNames))
	for _, name := range allowedNames {
		allowedNamesMap[name] = struct{}{}
	}
	return NameSampler{
		allow: allowedNamesMap,
	}
}

func (ns NameSampler) ShouldSample(parameters sdkTrace.SamplingParameters) sdkTrace.SamplingResult {
	psc := trace.SpanContextFromContext(parameters.ParentContext)

	if _, ok := ns.allow[parameters.Name]; ok {
		return sdkTrace.SamplingResult{
			Decision:   sdkTrace.RecordAndSample,
			Tracestate: psc.TraceState(),
		}
	}

	return sdkTrace.SamplingResult{
		Decision:   sdkTrace.Drop,
		Tracestate: psc.TraceState(),
	}
}

func (ns NameSampler) Description() string {
	return fmt.Sprintf("NameBased:{%v}", ns.allow)
}
