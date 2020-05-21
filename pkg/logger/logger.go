package logger

import (
	"log"
	"os"
)

//Severity type/level of logs required
type Severity struct {
	//Info, Warning, Error
	level string
}

//LogSeverity hold the data and severity level
type LogSeverity struct {
	level   Severity
	message string
}

//Logger take two paramters
//file and string
func Logger(data LogSeverity) {
	if data.level == "Info" {
		file, _ := os.OpenFile("../../go.mod")
	} else {
		log.Print("")
	}

}
