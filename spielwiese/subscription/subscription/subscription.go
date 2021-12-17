package subscription

import (
	"fmt"
	"ttslight/subscription/variable"
)

type Subscription struct {
	GroupTopic          string
	MonitoringFrequency int
	Qos                 int
	Variables           []variable.Variable
}

func New(groupTopic string, monitoringFrequency int, qos int) Subscription {
	e := Subscription{groupTopic, monitoringFrequency, qos, nil}
	return e
}

func (s *Subscription) AddVariables(vars []variable.Variable) {
	s.Variables = vars
}

func (s Subscription) ToString() string {
	var result string = fmt.Sprintf("[Subscription: topic: %v, monitoring frequency: %v, qos: %v, variables: '%v']", s.GroupTopic, s.MonitoringFrequency, s.Qos, s.Variables)
	return result
}
