package groupedsinks

import (
	"sync"

	"code.cloudfoundry.org/loggregator/doppler/internal/sinks"
	"code.cloudfoundry.org/loggregator/metricemitter"
	"github.com/cloudfoundry/sonde-go/events"
)

// MetricClient creates new CounterMetrics to be emitted periodically.
type MetricClient interface {
	NewCounter(name string, opts ...metricemitter.MetricOption) *metricemitter.Counter
}

type MetricBatcher interface {
	BatchIncrementCounter(name string)
}

func NewGroupedSinks(b MetricBatcher, mc MetricClient) *GroupedSinks {
	droppedMetric := mc.NewCounter("sinks.dropped",
		metricemitter.WithVersion(2, 0),
	)
	errorMetric := mc.NewCounter("sinks.errors.dropped",
		metricemitter.WithVersion(2, 0),
	)

	return &GroupedSinks{
		apps:          make(map[string]*AppGroup),
		batcher:       b,
		droppedMetric: droppedMetric,
		errorMetric:   errorMetric,
	}
}

type GroupedSinks struct {
	sync.RWMutex

	apps          map[string]*AppGroup
	batcher       MetricBatcher
	droppedMetric *metricemitter.Counter
	errorMetric   *metricemitter.Counter
}

func (group *GroupedSinks) RegisterAppSink(in chan<- *events.Envelope, sink sinks.Sink) bool {
	group.Lock()
	defer group.Unlock()

	appId := sink.AppID()
	if appId == "" || sink.Identifier() == "" {
		return false
	}

	sinksForApp, ok := group.apps[appId]
	if !ok || sinksForApp == nil {
		sinksForApp = NewAppGroup(
			group.batcher,
			group.droppedMetric,
			group.errorMetric,
		)
		group.apps[appId] = sinksForApp
	}
	return sinksForApp.AddSink(sink, in)
}

func (group *GroupedSinks) Broadcast(appId string, msg *events.Envelope) {
	group.RLock()
	defer group.RUnlock()

	sinksForApp, ok := group.apps[appId]
	if ok && sinksForApp != nil {
		sinksForApp.BroadcastMessage(msg)
	}
}

func (group *GroupedSinks) BroadcastError(appId string, msg *events.Envelope) {
	group.RLock()
	defer group.RUnlock()

	sinksForApp, ok := group.apps[appId]
	if ok && sinksForApp != nil {
		sinksForApp.BroadcastError(msg)
	}
}

func (group *GroupedSinks) DrainFor(appId, drainMetaData string) sinks.Sink {
	group.RLock()
	defer group.RUnlock()

	sinksForApp, ok := group.apps[appId]
	if !ok || sinksForApp == nil {
		return nil
	}
	return sinksForApp.Sink(drainMetaData)
}

func (group *GroupedSinks) DumpFor(appId string) *sinks.DumpSink {
	group.RLock()
	defer group.RUnlock()

	sinksForApp, ok := group.apps[appId]
	if !ok || sinksForApp == nil {
		return nil
	}
	return sinksForApp.RecentLogsSink(appId)
}

func (group *GroupedSinks) ContainerMetricsFor(appId string) *sinks.ContainerMetricSink {
	group.RLock()
	defer group.RUnlock()

	sinksForApp, ok := group.apps[appId]
	if !ok || sinksForApp == nil {
		return nil
	}
	return sinksForApp.ContainerMetricsSink("container-metrics-" + appId)
}

func (group *GroupedSinks) CloseAndDelete(sink sinks.Sink) bool {
	group.Lock()
	defer group.Unlock()

	appId := sink.AppID()

	sinksForApp, ok := group.apps[appId]
	if !ok || sinksForApp == nil {
		return false
	}

	removed := sinksForApp.RemoveSink(sink)
	if sinksForApp.IsEmpty() {
		delete(group.apps, appId)
	}

	return removed
}

func (group *GroupedSinks) DeleteAll() {
	group.Lock()
	defer group.Unlock()

	for appId, sinksForApp := range group.apps {
		if sinksForApp != nil {
			sinksForApp.RemoveAllSinks()
		}
		delete(group.apps, appId)
	}
}
