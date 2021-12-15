package main

import (
	"fmt"
	"io"
	"os"
	"ttslight/subscription/subscription"

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

func print_subscription() {
	defer logfile.Close()
	log.Info("Hello, World!")
	subscription1 := subscription.New("group_a", 5000, 1)
	log.WithFields(log.Fields{"topic": subscription1.GroupTopic}).Debug(subscription1.ToString())
	subscriptionLogger := log.WithFields(log.Fields{"topic": subscription1.GroupTopic})
	subscriptionLogger.Debug("Subscription Logger")
}

func main() {
	initLogging()
	print_subscription()
}
