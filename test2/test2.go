package test2

import "github.com/apex/log"

func String() bool {
	log.WithFields(log.Fields{
		"package": "test2",
	}).Info("info")
	return true
}
