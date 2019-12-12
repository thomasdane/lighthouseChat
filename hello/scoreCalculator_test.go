package main

import (
	"fmt"
	"testing"
)

func TestCalculateChange(t *testing.T) {

	for _, input := range inputs {

		//Arrange
		emojis := []string{"bad", "same", "good"}
		settings := UserSettings{Variance: input.variance, Emojis: emojis}

		//Act
		actual := CalculateChange(input.yesterday, input.today, settings)

		//Assert
		if actual != input.expected {
			error := fmt.Sprintf("For today %d yesterday %d variance %f, expected %s but got %s", input.today, input.yesterday, input.variance, input.expected, actual)
			t.Error(error)
		}
	}
}

type input struct {
	variance  float64
	yesterday int
	today     int
	expected  string
}

var inputs = []input{
	{2, 90, 93, "good"},
	{2, 90, 92, "same"},
	{2, 90, 90, "same"},
	{2, 90, 88, "same"},
	{2, 90, 87, "bad"},
	{3, 90, 87, "same"},
}
