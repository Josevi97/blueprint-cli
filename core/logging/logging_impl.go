package logging

import (
	"log"
	"os"
)

type LoggingImpl struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func (logging *LoggingImpl) write(logger log.Logger, msg string, args ...any) {
	logger.Printf(msg+"\n", args...)
}

func (logging *LoggingImpl) Info(msg string, args ...any) {
	logging.write(*logging.infoLog, msg, args...)
}

func (logging *LoggingImpl) Error(msg string, args ...any) {
	logging.write(*logging.errorLog, msg, args...)
}

func NewLogging(prefix string) Logging {
	infoLog := log.New(os.Stdout, "["+prefix+"] - ", log.Ldate|log.Ltime|log.Lmsgprefix)
	errorLog := log.New(os.Stderr, "", log.Ldate|log.Ltime)

	return &LoggingImpl{
		infoLog:  infoLog,
		errorLog: errorLog,
	}
}
