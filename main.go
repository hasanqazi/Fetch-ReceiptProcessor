package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"github.com/google/uuid"
)

// Receipt represents the JSON structure of the receipt payload
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
	GeneratedID  string `json:"-"`
	Points       int    `json:"-"`
}

// Item represents an item in the receipt
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// represents an in-memory storage for receipts
type InMemoryStorage struct {
	sync.RWMutex
	receipts map[string]Receipt
}

var storage InMemoryStorage

// ProcessReceiptHandler handles the /receipts/process endpoint
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the receipt
	receipt.GeneratedID = GenerateUniqueID()

	// Calculate the points for the receipt based on the rules
	receipt.Points = CalculatePoints(&receipt)

	// Store the receipt in the in-memory storage
	storage.Lock()
	storage.receipts[receipt.GeneratedID] = receipt
	storage.Unlock()

	// Send the response
	resp := struct {
		ID string `json:"id"`
	}{
		ID: receipt.GeneratedID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GetPointsHandler handles the /receipts/{id}/points endpoint
func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/receipts/")
	id = strings.TrimSuffix(id, "/points")

	storage.RLock()
	receipt, ok := storage.receipts[id]
	storage.RUnlock()

	if !ok {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	// Send the response
	resp := struct {
		Points int `json:"points"`
	}{
		Points: CalculatePoints(&receipt),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GenerateUniqueID generates a unique identifier for the receipt.
func GenerateUniqueID() string {
	return uuid.New().String()
}

func getAlphanumericLength(s string) int {
	count := 0
	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// CalculatePoints calculates the points for a given receipt based on the rules
func CalculatePoints(receipt *Receipt) int {
	points := 0
	println(points)
	// Calculate points based on retailer name
	points += getAlphanumericLength(receipt.Retailer)
	println("retailer name ", points)

	// Check if total is a round dollar amount with no cents
	totalFloat, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil && totalFloat == float64(int(totalFloat)) {
		points += 50
	}
	println("round dollar amount ", points)

	// Check if total is a multiple of 0.25
	if totalFloat != 0 && totalFloat/0.25 == float64(int(totalFloat/0.25)) {
		points += 25
	}
	println("total is a multiple of 0.25 ", points)

	// Calculate points based on the number of items
	points += len(receipt.Items) / 2 * 5
	println("number of items ", points)

	// Calculate points based on item description length
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price := parseFloat(item.Price)
			points += int(math.Ceil(price * 0.2)) // Round up to the nearest integer
		}
	}
	println("item description length ", points)

	// Check if day is odd
	purchaseDate := strings.Split(receipt.PurchaseDate, "/")
	if len(purchaseDate) >= 2 {
		day, _ := strconv.Atoi(purchaseDate[1])
		if day%2 == 1 {
			points += 6
		}
	}
	println("day is odd ", points)

	// Check if time is after 2:00pm and before 4:00pm
	purchaseTime := strings.Split(receipt.PurchaseTime, ":")
	if len(purchaseTime) >= 1 {
		hour, _ := strconv.Atoi(purchaseTime[0])
		if hour >= 14 && hour <= 16 {
			points += 10
		}
	}
	println("2:00pm and before 4:00pm ", points)

	return points
}

// parseFloat is a helper function to parse a string to a float64 value.
func parseFloat(s string) float64 {
	value, _ := strconv.ParseFloat(s, 64)
	return value
}

func main() {
	// Initialize the in-memory storage
	storage.receipts = make(map[string]Receipt)

	// Set up HTTP handlers
	http.HandleFunc("/receipts/process", ProcessReceiptHandler)
	http.HandleFunc("/receipts/", GetPointsHandler)

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
