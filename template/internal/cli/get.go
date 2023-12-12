package cli

import (
	"aocli/template/internal/aoc"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	year string
	day  string
	part int

	input bool

	auto bool
)

var cmdG = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Get a challenge, defaults to today (short: g)",
	Run: func(cmd *cobra.Command, args []string) {
		if input {
			in, err := aoc.GetInput(year, day)
			if err != nil {
				fmt.Println("🚨 An error occured:", err)
			}

			fmt.Println(in)
			return
		}

		chall, err := aoc.GetChallenge(year, day, part)
		if err != nil {
			fmt.Println("🚨 An error occured:", err)
		}

		if auto {
			fmt.Println(chall)
			return
		}

		fmt.Println("🎄 Getting challenge for day", day, "of year", year, ":")
		fmt.Printf("\n")
		fmt.Println(chall)
	},
}

func init() {
	defaultYear := aoc.CurrentYear()
	defaultDay := aoc.CurrentDay()

	cmdG.Flags().StringVarP(&year, "year", "y", defaultYear, "Specify the year for Advent of Code")
	cmdG.Flags().StringVarP(&day, "day", "d", defaultDay, "Specify the day for Advent of Code")
	cmdG.Flags().IntVarP(&part, "part", "p", 1, "Specify the part for Advent of Code")

	cmdG.Flags().BoolVarP(&auto, "auto", "a", false, "Output for use in scripts")
	cmdG.Flags().BoolVarP(&input, "input", "i", false, "Get the input for the challenge")

	rootCmd.AddCommand(cmdG)
}
