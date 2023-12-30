package pocketlog

import "io"

// Option defines a configuration function, an optional parameter to New
// that changes the behaviour of the Logger.
type Option func(*Logger)

// WithOutput returns a configuration function that sets the output of logs
func WithOutput(output io.Writer) Option {
    return func(lgr *Logger) {
        lgr.output = output
		}
}
