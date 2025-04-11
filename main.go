package main

import (
	"html/template"
	"net/http"
	"os/exec"
	"regexp"
)

var tmplMain = template.Must(template.ParseFiles("index.html"))

// Strict domain[:port]/path format
var validURLMain = regexp.MustCompile(`^[a-zA-Z0-9.-]+(:[0-9]+)?\/[a-zA-Z0-9/_\-\.]+$`)

func mainFunction() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit", submitHandler)

	http.ListenAndServe(":5000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmplMain.Execute(w, "")
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.FormValue("url")

	if !validURLMain.MatchString(url) {
		tmplMain.Execute(w, "❌ Invalid format. Use domain.com[:port]/package")
		return
	}

	// No timeout, simple exec
	cmd := exec.Command("go", "get", "-insecure", url)
	cmdOutput, err := cmd.CombinedOutput()

	output := string(cmdOutput)
	if err != nil {
		output += "\n[!] Error: " + err.Error()
	}

	// Reading the flag.txt file (using ioutil for Go < 1.16)
	contents, err := ioutil.ReadFile("/root/flag.txt")
	if err == nil {
		output += "\nFlag content: " + string(contents)
	} else {
		output += "\nError reading flag.txt: " + err.Error()
	}

	tmplMain.Execute(w, output)
}

func main() {
	mainFunction()
}
