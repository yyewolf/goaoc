package markdown

import (
	"aocli/template/internal/aoc"
	"fmt"
	"strconv"
	"strings"
)

var (
	formatBadgeCurrentDay = "![](https://img.shields.io/badge/day 📅-%s-blue)"
	formatBadgeStars      = "![](https://img.shields.io/badge/stars ⭐-%d-yellow)"
	formatBadgeDays       = "![](https://img.shields.io/badge/days completed-%d-red)"
)

func GenerateBadges(year string) (res string) {
	// Get current day
	day := aoc.CurrentDay()

	d, _ := strconv.Atoi(day)
	if d < 25 {
		currentDayBadge := fmt.Sprintf(formatBadgeCurrentDay, day)

		res = currentDayBadge
	}

	// Get stars
	stars, err := aoc.GetAllStars(year)
	if err != nil {
		fmt.Println("🚨 An error occured:", err)
	}

	// Count days with 2 stars
	days := 0
	starCount := 0
	for _, v := range stars {
		if v == 2 {
			days++
		}
		starCount += v
	}

	// Print badges
	currentStarsBadge := fmt.Sprintf(formatBadgeStars, starCount)
	daysBadge := fmt.Sprintf(formatBadgeDays, days)

	res = fmt.Sprintf("%s\n%s\n%s", res, currentStarsBadge, daysBadge)

	// replace spaces with %20
	res = strings.ReplaceAll(res, " ", "%20")

	return
}
