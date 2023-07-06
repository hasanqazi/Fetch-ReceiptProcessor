package main

import (
	"testing"
)

func TestGetAlphanumericLength(t *testing.T) {
	s := "Hello123"

	expectedLength := 8
	length := getAlphanumericLength(s)

	if length != expectedLength {
		t.Errorf("getAlphanumericLength returned %d, expected %d", length, expectedLength)
	}
}

func TestCalculatePoints(t *testing.T) {
	receipt := Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []Item{
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
		},
		Total: "9.00",
	}

	expectedPoints := 109
	points := CalculatePoints(&receipt)

	if points != expectedPoints {
		t.Errorf("CalculatePoints returned %d, expected %d", points, expectedPoints)
	}
}
