package util

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {

	var baseLogger = logrus.New()
	//baseLogger.SetReportCaller(true)

	var filename string = "../logfile.log"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	baseLogger.SetFormatter(&logrus.JSONFormatter{})

	if err != nil {
		// Cannot open log file. Logging to stderr
		fmt.Println(err)
	} else {
		baseLogger.SetOutput(f)
	}
	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PrettyPrint:     true,
	}

	return standardLogger
}

var (
	incommingRequest         = Event{1, "Incoming Request: %s"}
	incommingRequestWithBody = Event{2, "Incoming Body: %s "}
)

func (l *StandardLogger) IncomingRequest(requestUri string) {
	l.Infof(incommingRequest.message, requestUri)
}

func (l *StandardLogger) IncomingRequestWithBody(requestUri string, requestBody string) {
	l.Infof(incommingRequest.message, requestUri)
	l.Infof(incommingRequestWithBody.message, requestBody)
}
