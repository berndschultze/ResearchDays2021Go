package member

import log "github.com/sirupsen/logrus"

type member struct {
	name    string
	content int
}

var counter int

func New(name string, content int) member {
	e := member{name, content}
	return e
}

func (m member) PrintCounter() {
	log.Infof("Member %v Counter: %v", m.name, counter)
}

func (m member) SetCounter(count int) {
	log.Infof("Member %v Set Counter: %v", m.name, count)
	counter = count
}

func (m member) SetContent(cont int) {
	log.Infof("Member %v Set Content: %v", m.name, cont)
	m.content = cont
}

func (m member) PrintContent() {
	log.Infof("Member %v Content: %v", m.name, m.content)
}

func (m *member) SetContentRef(cont int) {
	log.Infof("Member %v Set Counter: %v", m.name, cont)
	m.content = cont
}

func (m *member) PrintContentRef() {
	log.Infof("Member %v Content: %v", m.name, m.content)
}
