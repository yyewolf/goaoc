package markdown

import (
	"fmt"
	"os"
	"regexp"
)

func GenerateYearsBut(year string) string {
	// List folders that match the pattern "20[0-9][0-9]" but are not the current years
	folders, err := os.ReadDir(".")
	if err != nil {
		return ""
	}

	re := regexp.MustCompile(`20[0-9][0-9]`)
	years := []string{}
	for _, folder := range folders {
		if folder.IsDir() && re.MatchString(folder.Name()) && folder.Name() != year {
			years = append(years, folder.Name())
		}
	}

	var output string

	if len(years) == 0 {
		return output
	}

	output += "### Other years\n\n"

	for _, year := range years {
		output += fmt.Sprintf("- [Advent Of Code %s](/%s)\n", year, year)
	}

	return output
}
