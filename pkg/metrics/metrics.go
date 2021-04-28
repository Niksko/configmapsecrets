package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

type ErrorType string
type SyncResult string

const (
	ErrorTypeKey              = "errorType"
	ErrorResourceKindKey      = "resourceKind"
	ErrorResourceNameKey      = "resourceName"
	ErrorResourceNamespaceKey = "resourceNamespace"
	SyncResultKey             = "result"

	GetError    ErrorType = "getError"
	DeleteError ErrorType = "deleteError"
	CreateError ErrorType = "createError"
	UpdateError ErrorType = "updateError"

	StatusError  ErrorType = "statusError"
	CleanupError ErrorType = "cleanupError"
	RenderError  ErrorType = "renderError"

	SyncSuccess = "success"
	SyncFailure = "failure"
)

var (
	syncs = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "configmapsecrets_syncs",
		Help: "Total number of syncs of the configmapsecrets controller",
	}, []string{SyncResultKey})

	internalErrors = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "configmapsecrets_internal_errors",
		Help: "Total number of internal errors of the configmapsecrets controller",
	}, []string{ErrorTypeKey, ErrorResourceKindKey, ErrorResourceNameKey, ErrorResourceNamespaceKey})

	ConfigMapSecretsMetrics = Metrics{
		internalErrors,
		syncs,
	}
)

type Metrics struct {
	internalErrors *prometheus.CounterVec
	syncs          *prometheus.CounterVec
}

func (m Metrics) IncrementTotalSyncs(syncResult SyncResult) {
	m.syncs.With(prometheus.Labels{SyncResultKey: string(syncResult)}).Inc()
}

func (m Metrics) IncrementInternalErrors(errorType ErrorType, kind, name, namespace string) {
	m.internalErrors.With(prometheus.Labels{
		ErrorTypeKey:              string(errorType),
		ErrorResourceKindKey:      kind,
		ErrorResourceNameKey:      name,
		ErrorResourceNamespaceKey: namespace,
	}).Inc()
}

func RegisterMetrics(registry metrics.RegistererGatherer) error {
	for _, reg := range []prometheus.Collector{syncs, internalErrors} {
		err := registry.Register(reg)
		if err != nil {
			return err
		}
	}
	return nil
}
