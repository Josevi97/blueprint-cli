package logging

import "log"

type Logging interface {
	// A hidden centralized method for logging
	write(logger log.Logger, msg string, args ...any)

	// Basic logging method to write info message in console
	Info(msg string, args ...any)

	// Logging method to write error message in console
	Error(msg string, args ...any)
}
