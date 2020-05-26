package logger

import (
	"log"
	"os"
)


// LogError in a file 
func LogError(err error) {
	if err != nil {
		//in production this should be changed in the configuration file
		path := "./s4.log"
		file, _ := os.OpenFile(path, os.O_CREATE | os.O_APPEND , 0644)
		defer file.Close()
		file.Write([]byte(err.Error()))
		log.Println(err)
	}
}