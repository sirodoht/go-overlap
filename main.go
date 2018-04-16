package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func main() {
	timezones := os.Args[1:]
	timezonesOffsets := []int{}
	for _, t := range timezones {
		if t[:3] != "utc" && t[:3] != "gmt" {
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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(timezones)
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	data := [][]string{}

	for _, tz := range timezonesOffsets {
		for i, v := range getOffset(tz) {
			if len(data) >= i {
				data = append(data, []string{})
			}
			timeStr := strconv.Itoa(v) + ":00"
			timeStrZeroPad := tablewriter.PadLeft(timeStr, "0", 5)
			timeStrSpacePad := tablewriter.Pad(timeStrZeroPad, " ", 9)
			data[i] = append(data[i], timeStrSpacePad)
		}
	}
	table.AppendBulk(data)
	fmt.Println()
	table.Render()
}

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

func invalidInput(message string) {
	fmt.Println()
	fmt.Println(message)
	fmt.Println()
	fmt.Println("Usage: overlap [timezone]...")
	fmt.Println("e.g. overlap utc-4 utc+3")
	os.Exit(2)
}
