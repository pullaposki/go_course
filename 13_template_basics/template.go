package main

import (
	"html/template"
	"log"
	"net/http"
)

type Item struct {
	Type string
	Count int
	Price int
}

type ShoppingData struct {
	ShoppingList string
	Items []Item
}

func main() {
    tmpl, err := template.ParseFiles("layout.html")
    if err != nil {
        log.Fatalf("Error parsing template: %v", err)
    }

    // Register a new route in the default ServeMux
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Create a new instance of ShoppingData
        data := ShoppingData{
            ShoppingList: "Shopping List",
            // Initialize the Items field with a slice of Item
            Items: []Item{
                {Type: "Apple", Count: 10, Price: 100},
                {Type: "Banana", Count: 5, Price: 50},
                {Type: "Orange", Count: 3, Price: 30},
            },
        }
        // Execute the template with the data and write the output to w
        err := tmpl.Execute(w, data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    })

		// Start the server on port 3000
    log.Fatal(http.ListenAndServe(":3000", nil))
}
