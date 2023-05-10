package main

import log "github.com/sirupsen/logrus"

var a int

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	for i := 0; i <= 10; i++ {
		a += i * i
		log.WithFields(log.Fields{
			"i": i,
			"a": a,
		}).Info("значения переменных")

	}

}
