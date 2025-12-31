package handlers

import (
	"ascii-art-web/asciigenerator"
	"bytes"
	"html/template"
	"net/http"
	"strings"
)

// Data represents the structure of the data passed to the HTML template.
type Data struct {
	Output string
}

// AsciiArtHandler manages the POST request, validates input, and renders the result.
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Restrict access to POST method only
	if r.Method != http.MethodPost {
		ErrorHandler(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve user input and banner choice from the form
	input := r.FormValue("input")
	banner := r.FormValue("banner")

	// Validate the banner selection
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		ErrorHandler(w, "400 Bad Request: Invalid banner choice", http.StatusBadRequest)
		return
	}

	// Ensure the input field exists in the form submission
	_, ok := r.Form["input"]
	if !ok {
		ErrorHandler(w, "400 Bad Request: Missing input field", http.StatusBadRequest)
		return
	}

	// Parse the index.html template from the templates directory
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorHandler(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Invoke the core logic to generate ASCII art
	output, err := asciigenerator.GenerateAsciiArt(input, banner)
	if err != nil {
		if strings.HasPrefix(err.Error(), "400") {
			ErrorHandler(w, err.Error(), http.StatusBadRequest)
		} else {
			ErrorHandler(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Prepare data for template execution
	data := Data{
		Output: output,
	}

	// Buffer the output to catch any execution errors before writing to the response
	var buff bytes.Buffer
	err = tmpl.Execute(&buff, data)
	if err != nil {
		ErrorHandler(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write the final rendered HTML to the ResponseWriter
	w.Write(buff.Bytes())
}
