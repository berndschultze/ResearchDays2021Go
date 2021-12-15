package subscription

import "fmt"

type subscription struct {
	GroupTopic          string
	MonitoringFrequency int
	Qos                 int
}

func New(groupTopic string, monitoringFrequency int, qos int) subscription {
	e := subscription{groupTopic, monitoringFrequency, qos}
	return e
}

func (s subscription) ToString() string {
	var result string = fmt.Sprintf("[Subscription: topic: %v, monitoring frequency: %v, qos: %v]", s.GroupTopic, s.MonitoringFrequency, s.Qos)
	return result
}
