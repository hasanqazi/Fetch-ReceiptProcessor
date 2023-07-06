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
		Retailer:     "Example Retailer",
		PurchaseDate: "2023/07/06",
		PurchaseTime: "15:30",
		Items: []Item{
			{
				ShortDescription: "Item 1",
				Price:            "9.99",
			},
			{
				ShortDescription: "Item 2",
				Price:            "4.50",
			},
		},
		Total: "14.49",
	}

	expectedPoints := 39
	points := CalculatePoints(&receipt)

	if points != expectedPoints {
		t.Errorf("CalculatePoints returned %d, expected %d", points, expectedPoints)
	}
}
