package main

import (
	"html/template"
	"net/http"
	"strconv"
)

// Contact struct represents a contact information
type Contact struct {
	Id        string
	Firstname string
	Lastname  string
	Email     string
	Phone     string
	Address   string
	City      string
}

// ContactListData struct holds a list of contacts
type ContactListData struct {
	Contacts []Contact
}

// FormData struct holds form data
type FormData struct {
	FormName string
}

func main() {

	// Initialize an empty slice of Contact structs and an id counter
	contactList := make([]Contact, 0)
	id := 100

	// Parse the HTML templates
	list_template := template.Must(template.ParseFiles("./views/list.html"))
	form_template := template.Must(template.ParseFiles("./views/form.html"))
	details_template := template.Must(template.ParseFiles("./views/details.html"))

	// Define the handler for the root ("/") path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// If the request method is POST, create a new Contact from the form data
		if r.Method == http.MethodPost {
			contact := Contact{
				Id:        strconv.FormatInt(int64(id), 10),
				Firstname: r.FormValue("firstname"),
				Lastname:  r.FormValue("lastname"),
				Email:     r.FormValue("email"),
				Phone:     r.FormValue("phone"),
				Address:   r.FormValue("address"),
				City:      r.FormValue("city"),
			}
			// Increment the id counter and add the new Contact to the list
			id++
			contactList = append(contactList, contact)
		}

		// Create a ContactListData struct with the current list of contacts
		data := ContactListData{Contacts: contactList}

		// Execute the list template with the current data
		list_template.Execute(w, data)
	})

	// Define the handler for the "/form" path
	http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
		// Create a FormData struct with a message
		data := FormData{FormName: "Please enter a new Contact"}
		// Execute the form template with the current data
		form_template.Execute(w, data)
	})

	// Define the handler for the "/contact/" path
	http.HandleFunc("/contact/", func(w http.ResponseWriter, r *http.Request) {

		// Create a Contact struct with the current data
		var data Contact

		// Get the id from the URL
		temp_string := r.URL.String()

		// Get the last 3 characters of the URL
		temp_id := temp_string[len(temp_string)-3:]

		// Loop through the contactList and find the contact with the matching id
		for _, contact := range contactList {
			if contact.Id == temp_id {
				data = contact
			}
		}

		// Execute the details template with the current data
		details_template.Execute(w, data)
	})

	// Start the server on port 3000
	http.ListenAndServe(":3000", nil)
}