package main

import "math"

//UserSettings holds the user's preferences
type UserSettings struct {
	Variance float64
	Emojis   []string
}

//CalculateChange returns an emoji based on the change between lighthouse runs
func CalculateChange(yesterday int, today int, settings UserSettings) (emoji string) {

	dailyChange := float64(today - yesterday)

	absoluteChange := math.Abs(dailyChange)

	if absoluteChange <= settings.Variance || absoluteChange == 0 {
		return settings.Emojis[1] //Return the no change emoji
	}

	if dailyChange < 0 {
		return settings.Emojis[0] //Things go worse
	}

	return settings.Emojis[2] //Things got better
}
