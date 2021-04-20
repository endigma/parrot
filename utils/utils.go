package utils

import log "github.com/sirupsen/logrus"

func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
