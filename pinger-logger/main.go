package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var statusFilePath = flag.String("oStatus", "/opt/webapp/status", "path to file to write the latest http status")
var logFilePath = flag.String("log", "/opt/webapp/pinger-logger.log", "path to log file")
var serviceAdress = flag.String("address", "localhost:8888/status", "http address to ping")

func main() {
	flag.Parse()

	log.SetFlags(log.LUTC | log.Ldate | log.Ltime)
	f, err := os.OpenFile(*logFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("can't open nor create log file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	resp, err := http.Get("http://" + *serviceAdress)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	UpdateStatus(resp.StatusCode)
}

func UpdateStatus(status int) error {
	statusVerbose := "SUCCESS"
	if status != 200 {
		log.Printf("error during pinging %s: HTTP status %d\n", *serviceAdress, status)
		statusVerbose = "ERROR"
	}

	f, err := os.OpenFile(*statusFilePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("can't open nor create status file: %v", err)
	}
	f.WriteString(fmt.Sprintf("%s %s\n", statusVerbose, time.Now().UTC().Format("2006/01/02 03:04:05")))

	return nil
}
