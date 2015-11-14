package util

import (
	"testing"
)

func TestParseServiceID(t *testing.T) {
	sample := []struct {
		serviceID string
		name      string
		status    int
		id        string
	}{
		{"one", "one", 0, ""},
		{"one.1", "one", 1, ""},
		{"one.1.id", "one", 1, "id"},
	}

	for _, v := range sample {
		info, err := ParseServiceID(v.serviceID)
		if err != nil {
			t.Error(err)
		}
		if info.ID() != v.id {
			t.Errorf("expected %s got %s", info.ID(), v.id)
		}
		if info.Name() != v.name {
			t.Errorf("expected %s got %s", info.Name(), v.name)
		}
		if info.Status() != v.status {
			t.Errorf("expected %s got 5s", info.Status(), v.status)
		}
	}
}
