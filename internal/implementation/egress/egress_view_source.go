package egress

import "kafmesh-example/internal/definitions/egress"

var _ egress.Egress_ViewSource = &EndpointViewsource{}

// EndpointViewsource sinks customer egress endpoints into Kafka.
type EndpointViewsource struct {
}

// Sync outputs all known egress endpoints to kafka.
func (vs *EndpointViewsource) Sync(ctx egress.Egress_ViewSource_Context) error {
	return nil
}
