package aoc

import (
	"aocli/template/internal/config"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

func GetUserID() (string, error) {
	c := config.C

	r, err := http.NewRequest("GET", "https://adventofcode.com/settings", nil)
	if err != nil {
		return "", err
	}
	r.AddCookie(&http.Cookie{Name: "session", Value: c.Secrets.AocSession})

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Search for "ownerproof-"
	// fmt.Println(string(data))
	i := strings.Index(string(data), "ownerproof-")
	if i == -1 {
		return "", fmt.Errorf("userid not found")
	}

	i += 11

	data = data[i:]

	// Search for next -
	j := strings.Index(string(data), "-")
	if j == -1 {
		return "", fmt.Errorf("userid not found")
	}

	data = data[:j]

	return string(data), nil
}

func GetChallenge(year, day string, part int) (string, error) {
	c := config.C

	r, err := http.NewRequest("GET", "https://adventofcode.com/"+year+"/day/"+day, nil)
	if err != nil {
		return "", err
	}
	r.AddCookie(&http.Cookie{Name: "session", Value: c.Secrets.AocSession})

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	p := bluemonday.NewPolicy()
	sanitized := p.Sanitize(string(data))

	// Replace HTML characters to their ASCII equivalent
	sanitized = strings.ReplaceAll(sanitized, "&amp;", "&")
	sanitized = strings.ReplaceAll(sanitized, "&lt;", "<")
	sanitized = strings.ReplaceAll(sanitized, "&gt;", ">")
	sanitized = strings.ReplaceAll(sanitized, "&quot;", "\"")
	sanitized = strings.ReplaceAll(sanitized, "&#39;", "'")
	sanitized = strings.ReplaceAll(sanitized, "&#34;", "\"")

	if part == 1 {
		i := strings.Index(string(sanitized), "--- Day "+day+": ")
		if i == -1 {
			return "", fmt.Errorf("day not found")
		}
		sanitized = sanitized[i+1:]

		i = strings.Index(string(sanitized), "---")

		sanitized = sanitized[i+3:]

		j := strings.Index(string(sanitized), "To begin, get your puzzle input")
		if j == -1 {
			j = strings.Index(string(sanitized), "Your puzzle answer was")
		}
		sanitized = sanitized[:j-2]

		return sanitized, nil
	} else if part == 2 {
		i := strings.Index(string(sanitized), "--- Part Two ---")
		if i == -1 {
			return "", fmt.Errorf("part not found")
		}

		sanitized = sanitized[i+16:]

		j := strings.Index(string(sanitized), "To begin, get your puzzle input")
		if j == -1 {
			j = strings.Index(string(sanitized), "Your puzzle answer was")
		}
		sanitized = sanitized[:j-2]
	}

	return sanitized, nil
}

func GetInput(year, day string) (string, error) {
	c := config.C

	r, err := http.NewRequest("GET", "https://adventofcode.com/"+year+"/day/"+day+"/input", nil)
	if err != nil {
		return "", err
	}
	r.AddCookie(&http.Cookie{Name: "session", Value: c.Secrets.AocSession})

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func GetStars(year, day string) (int, error) {
	c := config.C

	r, err := http.NewRequest("GET", "https://adventofcode.com/"+year+"/", nil)
	if err != nil {
		return 0, err
	}
	r.AddCookie(&http.Cookie{Name: "session", Value: c.Secrets.AocSession})

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Search for calendar-day%d
	i := strings.Index(string(data), fmt.Sprintf("calendar-day%s ", day))
	if i == -1 {
		return 0, fmt.Errorf("day not found")
	}

	if strings.HasPrefix(string(data)[i:], fmt.Sprintf("calendar-day%s calendar-verycomplete", day)) {
		return 2, nil
	} else if strings.HasPrefix(string(data)[i:], fmt.Sprintf("calendar-day%s calendar-complete", day)) {
		return 1, nil
	}

	return 0, nil
}

func GetAllStars(year string) (stars []int, err error) {
	c := config.C

	r, err := http.NewRequest("GET", "https://adventofcode.com/"+year+"/", nil)
	if err != nil {
		return
	}
	r.AddCookie(&http.Cookie{Name: "session", Value: c.Secrets.AocSession})

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	for day := 1; day <= 25; day++ {

		// Search for calendar-day%d
		i := strings.Index(string(data), fmt.Sprintf("calendar-day%d ", day))
		if i == -1 {
			continue
		}

		if strings.HasPrefix(string(data[i:]), fmt.Sprintf("calendar-day%d calendar-verycomplete", day)) {
			stars = append(stars, 2)
		} else if strings.HasPrefix(string(data[i:]), fmt.Sprintf("calendar-day%d calendar-complete", day)) {
			stars = append(stars, 1)
		} else {
			stars = append(stars, 0)
		}
	}

	return
}
