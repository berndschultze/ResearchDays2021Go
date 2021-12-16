package publication

import (
	"ttslight/subscription/subscription"
	"ttslight/subscription/variable"

	log "github.com/sirupsen/logrus"
)

type publication struct {
	GroupTopic          string
	MonitoringFrequency int
	Variables           []variable.Variable
}

func New(subs subscription.Subscription) publication {
	e := publication{subs.GroupTopic, subs.MonitoringFrequency, subs.Variables}
	return e
}

func (p publication) Publish() {
	log.Debugf("Publish publication for %v with variables %v", p.GroupTopic, p.Variables)
}
