package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Item struct {
	_Id string `json:"_id"`
	Type string `json:"type"`
	Count string `json:"count"`
	Price string `json:"price"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Session struct {
	TTL int64
	Token string
}

type MyToken struct {
	Token string `json:"token"`
}

type BackendMessage struct {
	Message string `json:"message"`
}

const time_to_live = 3600
var ShoppingItems []Item
var RegisteredUsers []User
var LoggedSessions []Session
var id int64
type Middleware func(http.HandlerFunc) http.HandlerFunc
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func HandleGetAndPost(resWriter http.ResponseWriter, req *http.Request) {
	switch req.Method {
		// Get all items
		case http.MethodGet:
			json.NewEncoder(resWriter).Encode(ShoppingItems)

			// Add a new item
		case http.MethodPost:
			var newItem Item

			// Decode the request body into the new item
			json.NewDecoder(req.Body).Decode(&newItem)

			// Assign an id to the new item
			newItem._Id = strconv.FormatInt(id, 10)
			id++

			// Append the new item to the list of items
			ShoppingItems = append(ShoppingItems, newItem)
			message := BackendMessage{Message: "Item added successfully"}

			// Encode the message and send it as a response
			json.NewEncoder(resWriter).Encode(message)
		default:
			resWriter.WriteHeader(http.StatusMethodNotAllowed)
			message := BackendMessage{Message: "Method not allowed"}
			json.NewEncoder(resWriter).Encode(message)
	}
}

func HandleDeleteAndPut(resWriter http.ResponseWriter, req *http.Request){
    temp_string := req.URL.String()
    temp_id := temp_string[len(temp_string)-3:]

    switch req.Method {
        case http.MethodDelete:
            for index, item := range ShoppingItems {
                if item._Id == temp_id {
                    ShoppingItems = append(ShoppingItems[:index], ShoppingItems[index+1:]...)
                    message := BackendMessage{Message: "Item deleted successfully"}
                    json.NewEncoder(resWriter).Encode(message)
                    return
                }
            }
        case http.MethodPut:
            var updatedItem Item

						// Decode the request body into the updated item
            json.NewDecoder(req.Body).Decode(&updatedItem)

						// Assign the id to the updated item
            for index, item := range ShoppingItems {
                if item._Id == temp_id {
                    ShoppingItems[index] = updatedItem
                    message := BackendMessage{Message: "Item updated successfully"}
                    json.NewEncoder(resWriter).Encode(message)
                    return
                }
            }
        default:
            resWriter.WriteHeader(http.StatusMethodNotAllowed)
            message := BackendMessage{Message: "Unknown method"}
            json.NewEncoder(resWriter).Encode(message)
    }
}

func CreateToken() string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create a slice of runes with a length of 64
	b := make([]rune, 64)

	// Generate a random token
	for i := range b {
		// Generate a random number between 0 and the length of the letters slice
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func Register(resWriter http.ResponseWriter, req *http.Request) {
	switch req.Method{
		case http.MethodPost:
			var user User
			json.NewDecoder(req.Body).Decode(&user)

			for _, temp_user := range RegisteredUsers {
				if user.Username == temp_user.Username {
					resWriter.WriteHeader(http.StatusConflict)
					message	:= BackendMessage{Message: "User already exists"}
					json.NewEncoder(resWriter).Encode(message)
					return
			}

			RegisteredUsers = append(RegisteredUsers, user)
			message := BackendMessage{Message: "User registered successfully"}
			json.NewEncoder(resWriter).Encode(message)

		}
		default:
			resWriter.WriteHeader(http.StatusMethodNotAllowed)
			message := BackendMessage{Message: "Method not allowed"}
			json.NewEncoder(resWriter).Encode(message)
		}
}

func main() {
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running on port 8080")
}