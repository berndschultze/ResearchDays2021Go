package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	varconfig "ttslight/config/variables"
	"ttslight/publication/publication"
	"ttslight/subscription/subscription"
	"ttslight/subscription/variable"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var logfile os.File

var subscriptions []*subscription.Subscription
var publications []*publication.Publication
var wg sync.WaitGroup

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

func createSubscriptions() {
	subscription1 := subscription.New("group_a", 2000, 1)
	subscription1.AddVariables(loadVariables())
	log.Debugf("Subscription with variables %v", subscription1.ToString())
	subscription2 := subscription.New("group_b", 1000, 1)
	subscription2.AddVariables(loadVariables())
	log.Debugf("Subscription with variables %v", subscription2.ToString())

	subscriptions = append(subscriptions, &subscription1, &subscription2)
}

func loadVariables() []*variable.Variable {
	data, err := varconfig.LoadVariables()
	if err != nil {
		log.Debugf("Error loading variable config %v", err)
	}
	var variables []*variable.Variable
	for _, varcon := range data {
		vari := variable.NewFromConf(&varcon)
		variables = append(variables, &vari)
	}
	return variables
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	log.Info("Endpoint Hit: homePage")
}

func returnSubscriptions(w http.ResponseWriter, r *http.Request) {
	log.Info("Endpoint Hit: returnSubscriptions")
	json.NewEncoder(w).Encode(subscriptions)
}

func returnSingleSubscription(w http.ResponseWriter, r *http.Request) {
	log.Info("Endpoint Hit: single subscription get")
	vars := mux.Vars(r)
	key := vars["group"]

	for _, sub := range subscriptions {
		if sub.GroupTopic == key {
			json.NewEncoder(w).Encode(sub)
		}
	}
}

func enableSubscription(w http.ResponseWriter, r *http.Request) {
	log.Info("Endpoint Hit: enable subscription")
	vars := mux.Vars(r)
	id := vars["group"]

	for _, pub := range publications {
		if pub.GroupTopic == id {
			log.Infof("Duplicate publication %v", pub.GroupTopic)
			http.Error(w, "409 conflict.", http.StatusConflict)
			return
		}
	}
	for _, sub := range subscriptions {
		if sub.GroupTopic == id {
			log.Infof("New Publication %v", sub.GroupTopic)
			pub := publication.New(sub)
			publications = append(publications, &pub)
			wg.Add(1)
			go pub.StartPublishing(&wg)
			return
		}
	}
	http.Error(w, "404 not found.", http.StatusNotFound)
}

func disableSubscription(w http.ResponseWriter, r *http.Request) {
	log.Info("Endpoint Hit: disable subscription")
	vars := mux.Vars(r)
	id := vars["group"]

	for index, pub := range publications {
		if pub.GroupTopic == id {
			pub.Stop()

			// Remove by overwrite with last and then cut out last, fast but no ordering preserved
			//publications[index] = publications[len(publications)-1]
			//publications = publications[:len(publications)-1]

			// Remove with keeping order
			publications = append(publications[:index], publications[index+1:]...)
			log.Infof("Removed publication %v", pub.GroupTopic)
			return
		}
	}
	http.Error(w, "404 not found.", http.StatusNotFound)
}

func runServer() {
	log.Info("Start Server")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnSubscriptions)
	myRouter.HandleFunc("/subscription/{group}", returnSingleSubscription).Methods(http.MethodGet)
	myRouter.HandleFunc("/subscription/{group}", enableSubscription).Methods(http.MethodPost)
	myRouter.HandleFunc("/subscription/{group}", disableSubscription).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	defer logfile.Close()
	initLogging()
	createSubscriptions()
	runServer()
}
