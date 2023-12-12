package aoc

import (
	"aocli/template/internal/config"
	"fmt"
	"time"
)

func CurrentYear() string {
	var currentYear = time.Now().Year()
	var year = fmt.Sprintf("%d", currentYear)

	if config.C.Public.CurrentYear != "" {
		year = config.C.Public.CurrentYear
	}

	return year
}

func CurrentDay() string {
	var currentDay = time.Now().Day()
	var day = fmt.Sprintf("%d", currentDay)

	return day
}
