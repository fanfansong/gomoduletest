package test1

import "github.com/apex/log"

// func String() bool {
// 	log.WithFields(log.Fields{
// 		"package": "test1",
// 	}).Info("info")
// 	return true
// }

func String2() {
	log.WithFields(log.Fields{
		"package":  "test1",
		"function": "String2()",
	}).Info("info")
	return true
}
