package test3

import "github.com/apex/log"

func String() bool {
	log.WithFields(log.Fields{
		"package": "test3",
	}).Info("info")
	return true
}
