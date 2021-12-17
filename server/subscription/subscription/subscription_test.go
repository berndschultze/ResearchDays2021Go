package subscription

import (
	"strings"
	"testing"
)

func TestInit(t *testing.T) {
	New("group_a", 5000, 1)
}

func TestToString(t *testing.T) {
	sub := New("group_a", 5000, 1)
	substr := sub.ToString()
	if !strings.Contains(substr, "group_a") {
		t.Error()
	}
}

func TestToStringTable(t *testing.T) {
	tables := []struct {
		sub                 Subscription
		groupTopic          string
		monitoringFrequency int
		qos                 int
	}{
		{New("group_a", 5000, 1), "group_a", 5000, 1},
		{New("group_b", 6000, 0), "group_b", 6000, 0},
		{New("group_c", 7000, 2), "group_c", 7000, 2},
		{New("group_d", 1000, 1), "group_d", 1000, 1},
	}

	for _, table := range tables {
		sub := New(table.groupTopic, table.monitoringFrequency, table.qos)
		if sub.ToString() != table.sub.ToString() {
			t.Errorf("To String Method not identical results, '%v' '%v'", sub.ToString(), table.sub.ToString())
		}
	}
}
