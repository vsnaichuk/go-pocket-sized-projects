package main

import (
	"os"
	"time"
  "learngo-pockets/pocketlog"
)

func main() {
		// Instantiating a new logger and calling a few methods to showcase usage
    lgr := pocketlog.New(pocketlog.LevelInfo, os.Stdout)
		lgr.Infof("A little copying is better than a little dependency.")
    lgr.Errorf("Errors are values. Documentation is for %s.", "users")
    lgr.Debugf("Make the zero (%d) value useful.", 0)
    lgr.Infof("Hallo, %d %v", 2024, time.Now())
}
