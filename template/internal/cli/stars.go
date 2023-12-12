package cli

import (
	"aocli/template/internal/aoc"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	allStars bool
)

func getSingleStars() {
	stars, err := aoc.GetStars(year, day)
	if err != nil {
		fmt.Println("🚨 An error occured:", err)
	}

	if auto {
		fmt.Println(stars)
		return
	}

	switch stars {
	case 0:
		// sad
		fmt.Println("🎄 No stars for day", day, "of year", year, ":(")
	case 1:
		// one star emoji
		fmt.Printf("🎄 You have : 🌟 (%s/%s)\n", year, day)
	case 2:
		// two star emoji
		fmt.Printf("🎄 You have : 🌟🌟 (%s/%s)\n", year, day)
	}
}

func getAllStars() {
	stars, err := aoc.GetAllStars(year)
	if err != nil {
		fmt.Println("🚨 An error occured:", err)
	}

	if auto {
		fmt.Println(stars)
		return
	}

	fmt.Println("🎄 Here are your stars for", year, ":")
	// Nice looking table going from first day to last day

	for i, c := range stars {
		s := ""
		for j := 0; j < c; j++ {
			s += "🌟"
		}

		fmt.Printf("%2d: %s\n", i+1, s)
	}

}

var cmdS = &cobra.Command{
	Use:     "stars",
	Aliases: []string{"s"},
	Short:   "Get stars, defaults to today (short: s)",
	Run: func(cmd *cobra.Command, args []string) {
		if !allStars {
			getSingleStars()
			return
		}
		getAllStars()
	},
}

func init() {
	defaultYear := aoc.CurrentYear()
	defaultDay := aoc.CurrentDay()

	cmdS.Flags().StringVarP(&year, "year", "y", defaultYear, "Specify the year for Advent of Code")
	cmdS.Flags().StringVarP(&day, "day", "d", defaultDay, "Specify the day for Advent of Code")

	cmdS.Flags().BoolVarP(&auto, "auto", "a", false, "Output for use in scripts")

	cmdS.Flags().BoolVarP(&allStars, "all", "A", false, "Get all stars for the current year")

	rootCmd.AddCommand(cmdS)
}
