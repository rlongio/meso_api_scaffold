// Package cap_errors is a command line utility for gathering NOAA product error information.
package main

import (
	"flag"
	"log"
	"strings"
	"time"
)

func main() {

	_ = flag.String("path", "", "root path of error files")
	_ = flag.String("log_path", "", "path of the log file")

	_ = flag.Uint("limit", 0, "limit the number of results to return")
	_ = flag.Bool("diagnostic", false, "Set to true for use as Debugger Diagnostic")

	flag.Parse()

	log.Printf("finished")
}

// toSlice converts string to a slice using "," as delimiter
func toSlice(value string) []string {
	return strings.Split(value, ", ")
}

// toDate takes a date string in the form of YYYY-MM-DD and
// returns a time.Time instnace
func toDate(date string) (t time.Time) {
	if date == "" {
		return time.Time{}
	}
	layout := "2006-01-02"
	t, err := time.ParseInLocation(layout, date, time.Local)
	if err != nil {
		panic(err)
	}
	return t
}
