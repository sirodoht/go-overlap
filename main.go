package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func main() {
	// check if help flag is passed
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "help" {
		usage()
		os.Exit(0)
	}

	// get timezones and their offsets from cli args
	timezones := os.Args[1:]
	timezonesOffsets := []int{}
	for _, t := range timezones {
		if len(t) < 3 || (t[:3] != "utc" && t[:3] != "gmt") {
			invalidInput("Invalid timezone: " + t)
		}
		tNum := t[3:]
		if tNum == "" {
			tNum = "0"
		}
		offset, err := strconv.Atoi(tNum)
		if err != nil {
			invalidInput("Invalid timezone: " + t)
		}
		if offset < -12 || offset > 12 {
			invalidInput("Invalid timezone: " + t)
		}
		timezonesOffsets = append(timezonesOffsets, offset)
	}

	// create ansii table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(timezones)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	// add times on data array
	data := [][]string{}
	for _, tz := range timezonesOffsets {
		for i, v := range getOffset(tz) {
			if len(data) >= i {
				data = append(data, []string{})
			}
			timeStr := strconv.Itoa(v) + ":00"
			timeStrZeroPad := tablewriter.PadLeft(timeStr, "0", 5)
			if v >= 9 && v <= 18 {
				timeStrZeroPad = "● " + timeStrZeroPad
			} else {
				timeStrZeroPad = "  " + timeStrZeroPad
			}
			timeStrSpacePad := tablewriter.Pad(timeStrZeroPad, " ", 12)
			data[i] = append(data[i], timeStrSpacePad)
		}
	}
	table.AppendBulk(data)

	// print result
	fmt.Println()
	table.Render()
}

// Get times by give offset.
func getOffset(offset int) []int {
	offsetTimes := []int{}
	if offset >= 0 {
		for i := offset; i < 24; i++ {
			offsetTimes = append(offsetTimes, i)
		}
		for i := 0; i < offset; i++ {
			offsetTimes = append(offsetTimes, i)
		}
	} else {
		for i := 24 + offset; i < 24; i++ {
			offsetTimes = append(offsetTimes, i)
		}
		for i := 0; i < 24+offset; i++ {
			offsetTimes = append(offsetTimes, i)
		}
	}
	return offsetTimes
}

// Print error message on invalid input along with usage details.
func invalidInput(errorMessage string) {
	fmt.Println()
	fmt.Println(errorMessage)
	usage()
	os.Exit(2)
}

// Print usage details.
func usage() {
	fmt.Println()
	fmt.Println("Usage: overlap [timezone]...")
	fmt.Println("e.g. overlap utc-4 utc+3")
}
