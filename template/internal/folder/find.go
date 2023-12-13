package folder

import (
	"os"
	"strings"
)

func FindRoot() string {
	var at = os.Getenv("PWD")

	for {
		if _, err := os.Stat(at + "/.git"); os.IsNotExist(err) {
			at = strings.TrimSuffix(at, "/"+strings.Split(at, "/")[len(strings.Split(at, "/"))-1])
		} else {
			break
		}
	}

	return at
}
