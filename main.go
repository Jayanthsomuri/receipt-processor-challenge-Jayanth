package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []Item  `json:"items"`
	Total        string  `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price           string `json:"price"`
}

type ReceiptID struct {
	ID string `json:"id"`
}

type PointsResponse struct {
	Points int64 `json:"points"`
}

var receipts = make(map[string]Receipt)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/receipts/process", processReceipt).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", getPoints).Methods("GET")
	fmt.Println("Server is running on port 8080...") 
	http.ListenAndServe(":8080", r)
}

func processReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid receipt", http.StatusBadRequest)
		return
	}

	// Validate receipt fields
	if receipt.Retailer == "" || receipt.PurchaseDate == "" || receipt.PurchaseTime == "" || len(receipt.Items) == 0 || receipt.Total == "" {
		http.Error(w, "Invalid receipt", http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the receipt
	id := uuid.New().String()
	receipts[id] = receipt

	// Return the ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ReceiptID{ID: id})
}

func getPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	receipt, exists := receipts[id]
	if !exists {
		http.Error(w, "No receipt found for that ID", http.StatusNotFound)
		return
	}

	points := calculatePoints(receipt)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PointsResponse{Points: points})
}

func calculatePoints(receipt Receipt) int64 {
	points := int64(0)

	// Rule 1: One point for every alphanumeric character in the retailer name.
	points += int64(len(regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(receipt.Retailer, "")))

	// Rule 2: 50 points if the total is a round dollar amount with no cents.
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int64(total)) {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25.
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	points += int64(len(receipt.Items) / 2 * 5)

	// Rule 5: If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer.
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {  
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int64(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd.
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}