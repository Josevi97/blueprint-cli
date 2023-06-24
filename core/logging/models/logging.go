package Logging

import "log"

type Logging interface {
	// A hidden centralized method for logging
	write(logger log.Logger, msg string, args ...any)

	// CLI standard message output
	Output(msg string, args ...any)

	// Basic logging method to write log info message in console
	Info(msg string, args ...any)

	// Logging method to write log error message in console
	Error(msg string, args ...any)
}
