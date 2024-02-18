package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type LineItem struct {
	ItemID   string `json:"item_id"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

type Order struct {
	CustomerID string     `json:"customer_id"`
	LineItems  []LineItem `json:"line_items"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	itemIDs := make([]string, 1000)
	for i := range itemIDs {
		itemIDs[i] = uuid.New().String()
	}

	customers := make([]string, 100)
	for i := range customers {
		customers[i] = uuid.New().String()
	}

	for i := 0; i < 120; i++ {
		customer := customers[rand.Intn(len(customers))]

		numLineItems := rand.Intn(10) + 1

		lineItems := make([]LineItem, numLineItems)
		for j := 0; j < numLineItems; j++ {
			itemID := itemIDs[rand.Intn(len(itemIDs))]
			lineItems[j] = LineItem{
				ItemID:   itemID,
				Quantity: rand.Intn(10) + 1,
				Price:    rand.Intn(10000) + 1,
			}
		}

		order := Order{
			CustomerID: customer,
			LineItems:  lineItems,
		}

		// Send the order
		jsonData, err := json.Marshal(order)
		if err != nil {
			panic(err)
		}
		_, err = http.Post("http://localhost:3000/orders", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			panic(err)
		}
		println("posted order", i+1)
	}
}
