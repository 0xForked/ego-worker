package helper

import (
	"fmt"
	"log"
)

func CheckError(err error, msg string) {
	// check error
	if err != nil {
		// show if error
		log.Fatalf("%s: %s", msg, err)
	}
}

func ShowMessage(message string) {
	// load app conf
	conf := GetConfig()
	// validate
	if conf.App.Debug {
		// show if valid
		fmt.Printf("%s \n", message)
	}
}
