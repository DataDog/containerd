package plugin

import (
	"fmt"
	"slices"

	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

type attributeSampler struct {
	selectedAttrs map[string][]string
}

// AttributeBased returns a Sampler that samples every trace matching a set of attributes.
// It only need to find one matching key-value pairs present in the attributes of the span.
func AttributeBased(selectedAttrs map[string][]string) attributeSampler {
	return attributeSampler{
		selectedAttrs: selectedAttrs,
	}
}

func (ss attributeSampler) ShouldSample(parameters sdkTrace.SamplingParameters) sdkTrace.SamplingResult {
	psc := trace.SpanContextFromContext(parameters.ParentContext)
	println(parameters.Name)
	for _, attr := range parameters.Attributes {
		println(attr.Key, attr.Value.Emit())
		if selectedValues, ok := ss.selectedAttrs[string(attr.Key)]; ok {
			if slices.Contains(selectedValues, attr.Value.Emit()) {
				println("Sampled")
				return sdkTrace.SamplingResult{
					Decision:   sdkTrace.RecordAndSample,
					Tracestate: psc.TraceState(),
				}
			}
		}
	}
	return sdkTrace.SamplingResult{
		Decision:   sdkTrace.Drop,
		Tracestate: psc.TraceState(),
	}
}

func (ss attributeSampler) Description() string {
	return fmt.Sprintf("AttributeBased:{%v}", ss.selectedAttrs)
}
