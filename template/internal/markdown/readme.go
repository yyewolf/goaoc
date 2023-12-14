package markdown

import (
	"strings"
	"text/template"
)

var readmeTemplate = template.Must(template.ParseFiles(mainTemplate))

type ReadmeData struct {
	Badges     string
	Year       string
	Stars      string
	Benchmarks string
	OtherYears string
}

func GenerateReadme(year string) string {
	w := strings.Builder{}

	var data ReadmeData

	data.Badges = GenerateBadges(year)
	data.Year = year
	data.Stars = GenerateStars(year)
	data.Benchmarks = GenerateBenches(year)
	data.OtherYears = GenerateYearsBut(year)

	readmeTemplate.ExecuteTemplate(&w, mainTemplate, data)

	return w.String()
}
