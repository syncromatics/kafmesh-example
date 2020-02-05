// Code generated by kafmesh-gen. DO NOT EDIT.

package definitions

import (
	"context"
	"time"

	"github.com/syncromatics/kafmesh/pkg/runner"
)

var (
	topics = []runner.Topic{
		runner.Topic {
			Name:       "kafmesh.deviceId.customer",
			Partitions: 10,
			Replicas:   1,
			Compact:    false,
			Retention:  86400000 * time.Millisecond,
			Segment:    43200000 * time.Millisecond,
			Create:     true,
		},
		runner.Topic {
			Name:       "kafmesh.customerId.details",
			Partitions: 10,
			Replicas:   1,
			Compact:    false,
			Retention:  86400000 * time.Millisecond,
			Segment:    43200000 * time.Millisecond,
			Create:     true,
		},
		runner.Topic {
			Name:       "kafmesh.deviceId.details",
			Partitions: 10,
			Replicas:   1,
			Compact:    false,
			Retention:  86400000 * time.Millisecond,
			Segment:    43200000 * time.Millisecond,
			Create:     true,
		},
		runner.Topic {
			Name:       "kafmesh.deviceId.enrichedDetails",
			Partitions: 10,
			Replicas:   1,
			Compact:    false,
			Retention:  86400000 * time.Millisecond,
			Segment:    43200000 * time.Millisecond,
			Create:     true,
		},
		runner.Topic {
			Name:       "kafmesh.deviceId.enrichedDetails-table",
			Partitions: 10,
			Replicas:   1,
			Compact:    false,
			Retention:  86400000 * time.Millisecond,
			Segment:    43200000 * time.Millisecond,
			Create:     true,
		},
		runner.Topic {
			Name:       "kafmesh.deviceId.heartbeat",
			Partitions: 10,
			Replicas:   1,
			Compact:    false,
			Retention:  86400000 * time.Millisecond,
			Segment:    43200000 * time.Millisecond,
			Create:     true,
		},
		runner.Topic {
			Name:       "kafmesh.deviceId.enrichedHeartbeat",
			Partitions: 10,
			Replicas:   1,
			Compact:    false,
			Retention:  86400000 * time.Millisecond,
			Segment:    43200000 * time.Millisecond,
			Create:     true,
		},
	}
)

func ConfigureTopics(ctx context.Context, brokers []string) error {
	return runner.ConfigureTopics(ctx, brokers, topics)
}