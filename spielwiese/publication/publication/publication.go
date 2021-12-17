package publication

import (
	"sync"
	"time"
	"ttslight/subscription/subscription"
	"ttslight/subscription/variable"

	log "github.com/sirupsen/logrus"
)

type publication struct {
	GroupTopic          string
	MonitoringFrequency int
	Variables           []variable.Variable
}

func New(subs *subscription.Subscription) publication {
	e := publication{subs.GroupTopic, subs.MonitoringFrequency, subs.Variables}
	return e
}

func (p *publication) Publish(index int) {
	log.Debugf("Publish publication number %v for %v with variables %v", index, p.GroupTopic, p.Variables)
}

func (p *publication) StartPublishing(wg *sync.WaitGroup, number int) {
	defer wg.Done()

	for i := 1; i <= number; i++ {
		p.Publish(i)
		time.Sleep(time.Duration(p.MonitoringFrequency) * time.Millisecond)
	}
}
