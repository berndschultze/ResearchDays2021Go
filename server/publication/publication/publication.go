package publication

import (
	"strings"
	"sync"
	"time"
	"ttslight/subscription/subscription"
	"ttslight/subscription/variable"

	log "github.com/sirupsen/logrus"
)

type Publication struct {
	GroupTopic          string
	MonitoringFrequency int
	Variables           []*variable.Variable
	running             bool
}

func New(subs *subscription.Subscription) Publication {
	e := Publication{subs.GroupTopic, subs.MonitoringFrequency, subs.Variables, false}
	return e
}

func (p *Publication) Publish(index int) {
	var varistr []string
	for _, vari := range p.Variables {
		varistr = append(varistr, vari.ToString())
	}
	log.Debugf("Publish publication number %v for %v with variables %v", index, p.GroupTopic, strings.Join(varistr, ", "))
}

func (p *Publication) StartPublishing(wg *sync.WaitGroup) {
	defer wg.Done()
	defer p.deferStop()

	p.running = true

	counter := 1

	for p.running {
		p.Publish(counter)
		counter++
		time.Sleep(time.Duration(p.MonitoringFrequency) * time.Millisecond)
	}
}

func (p *Publication) deferStop() {
	log.Infof("Defer Stop %v", p.GroupTopic)
	p.Stop()
}

func (p *Publication) Stop() {
	p.running = false
	log.Infof("Stopped %v", p.GroupTopic)
}
