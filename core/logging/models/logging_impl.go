package Logging

import (
	"log"
	"os"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorOrange = "\033[33m"
	colorCyan   = "\033[36m"
)

type LoggingImpl struct {
	outputLog *log.Logger
	infoLog   *log.Logger
	errorLog  *log.Logger
}

func (logging *LoggingImpl) write(logger log.Logger, msg string, args ...any) {
	// This should be the way
	// if (!initialized) return

	logger.Printf(msg+"\n", args...)
}

func (logging *LoggingImpl) Output(msg string, args ...any) {
	logging.write(*logging.outputLog, msg, args...)
}

func (logging *LoggingImpl) Info(msg string, args ...any) {
	logging.write(*logging.infoLog, colorCyan+msg+colorReset, args...)
}

func (logging *LoggingImpl) Error(msg string, args ...any) {
	logging.write(*logging.errorLog, colorRed+"Error: "+msg+colorReset, args...)
}

func NewLogging(prefix string) Logging {
	// Another way is to avoid the output if not initialized
	// if (!initialized) {
	// 	soutput := none
	// }

	outputLog := log.New(os.Stdout, "", 0)
	infoLog := log.New(os.Stdout, colorOrange+"["+prefix+"] "+colorReset, log.LstdFlags|log.Lmsgprefix)
	errorLog := log.New(os.Stderr, colorOrange+"["+prefix+"] "+colorReset, log.LstdFlags|log.Lmsgprefix|log.Llongfile)

	return &LoggingImpl{
		outputLog: outputLog,
		infoLog:   infoLog,
		errorLog:  errorLog,
	}
}
