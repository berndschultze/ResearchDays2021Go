package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

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

func getSubscriptions() {
	resp, err := http.Get("http://localhost:10000/all")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf("Code: %v, Body: '%v'", resp.StatusCode, sb)

}

func enablePublication(group string) {
	url := fmt.Sprintf("http://localhost:10000/subscription/%v", group)
	resp, err := http.Post(url, "text/plain", nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Code: %v", resp.StatusCode)
}

func disablePublication(group string) {
	url := fmt.Sprintf("http://localhost:10000/subscription/%v", group)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Code: %v", resp.StatusCode)
}

func main() {
	defer logfile.Close()
	initLogging()

	getSubscriptions()
	time.Sleep(2 * time.Second)
	enablePublication("group_a")
	enablePublication("group_a")
	time.Sleep(5 * time.Second)
	enablePublication("group_b")
	time.Sleep(5 * time.Second)
	disablePublication("group_a")
	disablePublication("group_a")
	time.Sleep(5 * time.Second)
	disablePublication("group_b")
}
