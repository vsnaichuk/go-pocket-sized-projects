package pocketlog

// Level represents one available logging level
type Level byte

const (
	/* LevelDebug represents the lowest level of log, mostly used for debugging purposes.
	 * "iota" allows us to create a sequence of numbers incremented on each line
	 */
	LevelDebug Level = iota

	/* LevelInfo represents a logging level
	 * that contains information deemed valuable.
	 */
	LevelInfo

	/* LevelError represents the highest logging level,
	 * only to be used to trace errors.
	 */
	LevelError
)
