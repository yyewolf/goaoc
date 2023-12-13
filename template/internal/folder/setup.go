package folder

import (
	"os"
	"os/exec"
)

func CreateDay(year, day string) {
	// Create ./year/day/
	// Copy what's in ./template/go ./year/day
	os.MkdirAll("./"+year+"/"+day, 0755)
	c := "cp -r ./template/go/* ./" + year + "/" + day
	cmd := exec.Command("sh", "-c", c)
	cmd.Run()
}
