package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tpl.ExecuteTemplate(w, "welcome.html", nil)
}


func ascii(w http.ResponseWriter, r *http.Request) {
	type Sub struct {
		Text   string
		Banner string
	}

	var s Sub

	s.Text = r.FormValue("text")
	s.Banner = r.FormValue("banner")

	if s.Text == "" || s.Banner == "" { // if there is no user input
		
	http.Error(w, " 400 Bad request", http.StatusBadRequest)
		return 
	}

	if s.Text[0] <= ' ' || s.Text[0] >= '~' { // less than first ascii char " " in .txt file, more than last char ~		http.Error(w," ",http.StatusBadRequest)
		http.Error(w, " 400 Bad request", http.StatusBadRequest)
		return
	}

	file, err := os.Open(s.Banner + ".txt") // opens the file
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return 

	}
	defer file.Close() // closes the file

	scanner := bufio.NewScanner(file) // reads the file
	// scanner.Split(bufio.ScanWords)

	var asciiString []string

	for scanner.Scan() {
		asciiString = append(asciiString, scanner.Text()) // appends the contents of the file to a empty string slice
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	myMap := make(map[int][]string)

	start := 31
	c := start

	// Iterate through each line
	for _, line := range asciiString { // for range through the lines of the txt file
		if line == "" {
			c += 1 // every time the loop encounters a empty line it adds 1 to variable c
		} else {
			myMap[c] = append(myMap[c], line)
		}
	}

	var lines string

	newline := strings.ReplaceAll(s.Text, "\r\n", "\\n")
	arrinput := strings.Split(newline, "\\n")
	for u := range arrinput {
		for y := 0; y < 8; y++ {
			for x := 0; x < len(arrinput[u]); x++ {
				ch := int(arrinput[u][x])

				lines += myMap[ch][y]
			}
			lines += "\n"

		}
	}
	tpl.ExecuteTemplate(w, "ascii-art.html", lines)
}
func main() {
	http.HandleFunc("/", welcome)
	fmt.Println("Starting the server on :8080...")
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	http.HandleFunc("/ascii-art", ascii)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
