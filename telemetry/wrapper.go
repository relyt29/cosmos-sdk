package telemetry

import (
	"time"

	"github.com/hashicorp/go-metrics"
)

// Common metric key constants
const (
	MetricKeyBeginBlocker       = "begin_blocker"
	MetricKeyEndBlocker         = "end_blocker"
	MetricKeyPrepareCheckStater = "prepare_check_stater"
	MetricKeyPrecommiter        = "precommiter"
	MetricLabelNameModule       = "module"
)

// GlobalLabels defines the set of global labels that will be applied to all
// metrics emitted using the telemetry package function wrappers.
var GlobalLabels = []metrics.Label{} // nolint: ignore // false positive

// NewLabel creates a new instance of Label with name and value
func NewLabel(name, value string) metrics.Label {
	return metrics.Label{Name: name, Value: value}
}

// ModuleMeasureSince provides a short hand method for emitting a time measure
// metric for a module with a given set of keys. If any global labels are defined,
// they will be added to the module label.
func ModuleMeasureSince(module string, start time.Time, keys ...string) {
	metrics.MeasureSinceWithLabels(
		keys,
		start.UTC(),
		append([]metrics.Label{NewLabel(MetricLabelNameModule, module)}, GlobalLabels...),
	)
}

// ModuleSetGauge provides a short hand method for emitting a gauge metric for a
// module with a given set of keys. If any global labels are defined, they will
// be added to the module label.
func ModuleSetGauge(module string, val float32, keys ...string) {
	metrics.SetGaugeWithLabels(
		keys,
		val,
		append([]metrics.Label{NewLabel(MetricLabelNameModule, module)}, GlobalLabels...),
	)
}

// IncrCounter provides a wrapper functionality for emitting a counter metric with
// global labels (if any).
func IncrCounter(val float32, keys ...string) {
	metrics.IncrCounterWithLabels(keys, val, GlobalLabels)
}

// IncrCounterWithLabels provides a wrapper functionality for emitting a counter
// metric with global labels (if any) along with the provided labels.
func IncrCounterWithLabels(keys []string, val float32, labels []metrics.Label) {
	metrics.IncrCounterWithLabels(keys, val, append(labels, GlobalLabels...))
}

// SetGauge provides a wrapper functionality for emitting a gauge metric with
// global labels (if any).
func SetGauge(val float32, keys ...string) {
	metrics.SetGaugeWithLabels(keys, val, GlobalLabels)
}

// SetGaugeWithLabels provides a wrapper functionality for emitting a gauge
// metric with global labels (if any) along with the provided labels.
func SetGaugeWithLabels(keys []string, val float32, labels []metrics.Label) {
	metrics.SetGaugeWithLabels(keys, val, append(labels, GlobalLabels...))
}

// MeasureSince provides a wrapper functionality for emitting a a time measure
// metric with global labels (if any).
func MeasureSince(start time.Time, keys ...string) {
	metrics.MeasureSinceWithLabels(keys, start.UTC(), GlobalLabels)
}
