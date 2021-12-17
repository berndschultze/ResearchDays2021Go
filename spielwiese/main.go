package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
	varconfig "ttslight/config/variables"
	"ttslight/member"
	"ttslight/publication/publication"
	"ttslight/subscription/subscription"
	"ttslight/subscription/variable"

	log "github.com/sirupsen/logrus"
)

var logfile os.File

func initLogging() {
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel)

	var logdirpath string = "log"
	var logfilepath string = fmt.Sprintf("%v/debug.log", logdirpath)

	os.Mkdir(logdirpath, 0755)
	os.Remove(logfilepath)

	logfile, err := os.OpenFile(logfilepath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Info("error opening file: %v", err)
	}

	mw := io.MultiWriter(os.Stdout, logfile)

	log.SetOutput(mw)
}

func printSubscription() {
	defer logfile.Close()
	log.Info("Hello, World!")
	subscription1 := subscription.New("group_a", 5000, 1)
	log.WithFields(log.Fields{"topic": subscription1.GroupTopic}).Debug(subscription1.ToString())
	subscriptionLogger := log.WithFields(log.Fields{"topic": subscription1.GroupTopic})
	subscriptionLogger.Debug("Subscription Logger")
	subscription1.AddVariables(loadVariables())
	log.Debugf("Subscription with variables %v", subscription1.ToString())
}

func printVariables() {
	data := loadVariables()
	for _, variable := range data {
		log.Infof("variable loaded from config %v", variable.ToString())
	}
}

func loadVariables() []variable.Variable {
	data, err := varconfig.LoadVariables()
	if err != nil {
		log.Debugf("Error loading variable config %v", err)
	}
	var variables []variable.Variable
	for _, varcon := range data {
		variables = append(variables, variable.NewFromConf(&varcon))
	}
	return variables
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return true
	case <-time.After(timeout):
		return false
	}
}

func publish() {
	subscription1 := subscription.New("group_a", 2000, 1)
	subscription1.AddVariables(loadVariables())
	publication1 := publication.New(&subscription1)
	publication1.Publish(0)
	log.Debug("Start Publish")
	var wg sync.WaitGroup
	wg.Add(1)
	go publication1.StartPublishing(&wg, 5)
	time.Sleep(2 * time.Second)
	log.Debug("Publishing")
	//wg.Wait()
	result := waitTimeout(&wg, 10*time.Second)
	log.Debugf("Wait Timeout: %v", result)
}

func membering() {
	mem1 := member.New("lala", 1)
	mem2 := member.New("lolo", 2)

	mem1.PrintCounter()
	mem2.PrintCounter()
	mem1.SetCounter(3)
	mem2.SetCounter(4)
	mem1.PrintCounter()
	mem2.PrintCounter()

	mem1.SetContent(5)
	mem2.SetContent(6)
	mem1.PrintContent()
	mem2.PrintContent()

	mem1.SetContentRef(50)
	mem2.SetContentRef(60)
	mem1.PrintContent()
	mem2.PrintContent()

	mem1.SetContentRef(500)
	mem2.SetContentRef(600)
	mem1.PrintContentRef()
	mem2.PrintContentRef()
}

func main() {
	initLogging()
	membering()
	printSubscription()
	printVariables()
	publish()
}
