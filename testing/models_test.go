package testing

import (
	"strconv"

	"github.com/golang/protobuf/ptypes"

	gherkin "github.com/cucumber/gherkin-go"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pkg/errors"
)

type Device struct {
	ID int64
}

func NewDevice(data *gherkin.DataTable) (*Device, error) {
	m := dataToMap(data)

	v, ok := m["id"]
	if !ok {
		return nil, errors.Errorf("id was not set and is required")
	}

	id, err := strconv.Atoi(v)
	if err != nil {
		return nil, err
	}

	return &Device{
		ID: int64(id),
	}, nil
}

type Customer struct {
	ID   int64
	Name *string
}

func NewCustomer(data *gherkin.DataTable) (*Customer, error) {
	m := dataToMap(data)

	v, ok := m["id"]
	if !ok {
		return nil, errors.Errorf("id was not set and is required")
	}

	id, err := strconv.Atoi(v)
	if err != nil {
		return nil, err
	}

	var name *string
	n, ok := m["name"]
	if ok {
		name = &n
	}

	return &Customer{
		ID:   int64(id),
		Name: name,
	}, nil
}

type Details struct {
	Name string
	Time *timestamp.Timestamp
}

func NewDetails(data *gherkin.DataTable) (*Details, error) {
	m := dataToMap(data)

	v, ok := m["name"]
	if !ok {
		return nil, errors.Errorf("name was required but not provided")
	}

	t, ok := m["time"]
	if !ok {
		return nil, errors.Errorf("time is required but not provided")
	}

	var time *timestamp.Timestamp
	switch t {
	case "now":
		time = ptypes.TimestampNow()
	default:
		return nil, errors.Errorf("unknown time '%s'", t)
	}

	return &Details{
		Name: v,
		Time: time,
	}, nil
}

type Heartbeat struct {
	Time      *timestamp.Timestamp
	IsHealthy bool
}

func NewHeartbeat(data *gherkin.DataTable) (*Heartbeat, error) {
	m := dataToMap(data)

	t, ok := m["time"]
	if !ok {
		return nil, errors.Errorf("time is required but not provided")
	}

	var time *timestamp.Timestamp
	switch t {
	case "now":
		time = ptypes.TimestampNow()
	default:
		return nil, errors.Errorf("unknown time '%s'", t)
	}

	h, ok := m["isHealthy"]
	if !ok {
		return nil, errors.Errorf("isHealthy is required but not provided")
	}

	healthy := false
	switch h {
	case "true":
		healthy = true
	case "false":
		healthy = false
	default:
		return nil, errors.Errorf("unknown value for isHealth '%s', must be'true' or 'false", h)
	}

	return &Heartbeat{
		Time:      time,
		IsHealthy: healthy,
	}, nil
}

func dataToMap(data *gherkin.DataTable) map[string]string {
	m := map[string]string{}
	for _, r := range data.Rows {
		m[r.Cells[0].Value] = r.Cells[1].Value
	}

	return m
}
