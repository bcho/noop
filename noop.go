// A program that do nothing but sleep.
package main

import (
	"flag"
	"log"
	"time"
)

const TICK_SECS = 10000

var (
	tickSecs int64
	verbose  bool
)

func main() {
	flag.Parse()

	for _ = range time.Tick(time.Duration(tickSecs) * time.Second) {
		if verbose {
			log.Printf("tick!")
		}
	}
}

func init() {
	flag.Int64Var(&tickSecs, "tick", TICK_SECS, "tick seconds")
	flag.BoolVar(&verbose, "verbose", false, "verbose mode")
}
