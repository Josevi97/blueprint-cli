package logging

import (
	"log"
	"os"
)

type LoggingImpl struct {
	outputLog *log.Logger
	infoLog   *log.Logger
	errorLog  *log.Logger
}

func (logging *LoggingImpl) write(logger log.Logger, msg string, args ...any) {
	logger.Printf(msg+"\n", args...)
}

func (logging *LoggingImpl) Output(msg string, args ...any) {
	logging.write(*logging.outputLog, msg, args...)
}

func (logging *LoggingImpl) Info(msg string, args ...any) {
	logging.write(*logging.infoLog, msg, args...)
}

func (logging *LoggingImpl) Error(msg string, args ...any) {
	logging.write(*logging.errorLog, msg, args...)
}

func NewLogging(prefix string) Logging {
	outputLog := log.New(os.Stdout, "", 0)
	infoLog := log.New(os.Stdout, "["+prefix+"] - ", log.Ldate|log.Ltime|log.Lmsgprefix)
	errorLog := log.New(os.Stderr, "["+prefix+"] - Error: ", log.Ldate|log.Ltime|log.Lmsgprefix|log.Llongfile)

	return &LoggingImpl{
		outputLog: outputLog,
		infoLog:   infoLog,
		errorLog:  errorLog,
	}
}
