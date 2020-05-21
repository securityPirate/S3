package logger

import (
	"log"
	"os"
)


// LogError in a file 
func LogError(err error) {
	if err != nil {
		file, _ := os.OpenFile("s4.log", os.O_CREATE | os.O_APPEND , 0644)
		defer file.Close()
		file.Write([]byte(err.Error()))
		log.Println(err)
	}
}