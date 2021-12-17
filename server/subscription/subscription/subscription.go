package subscription

import (
	"fmt"
	"strings"
	"ttslight/subscription/variable"
)

type Subscription struct {
	GroupTopic          string
	MonitoringFrequency int
	Qos                 int
	Variables           []*variable.Variable
}

func New(groupTopic string, monitoringFrequency int, qos int) Subscription {
	e := Subscription{groupTopic, monitoringFrequency, qos, nil}
	return e
}

func (s *Subscription) AddVariables(vars []*variable.Variable) {
	s.Variables = vars
}

func (s *Subscription) ToString() string {
	var varistr []string
	for _, vari := range s.Variables {
		varistr = append(varistr, vari.ToString())
	}
	var result string = fmt.Sprintf("[Subscription: topic: %v, monitoring frequency: %v, qos: %v, variables: '%v']", s.GroupTopic, s.MonitoringFrequency, s.Qos, strings.Join(varistr, ", "))
	return result
}
