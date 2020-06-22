package main

import "log"

func main() {
	Infof("process_name", "Start.")
	Errorf("process_name", "Error occurred.")
	Infof("process_name", "End.")
}

func Infof(tag string, message string) {
	log.SetPrefix("[INFO]")
	Logf(tag, message)
}

func Errorf(tag string, message string) {
	log.SetPrefix("[Error]")
	Logf(tag, message)
}

func Logf(tag string, message string) {
	log.Printf("%s, %s\n", tag, message)
}
