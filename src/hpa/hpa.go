package main

import (
	
	"fmt"
	"math"
	"html/template"
	"net/http"
)

var x = 0.0001

func Looping() (r string) {
	for i := 0; i < 1000000; i++ {
		x += math.Sqrt(x)
	}
	
	return "Code.education Rocks!"
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("index.html")
	msg := Looping()
	w.WriteHeader(http.StatusOK)
	tpl.ExecuteTemplate(w, "index.html", template.HTML(msg))
}

func main() {
	
	//fmt.Println(Looping())
	http.HandleFunc("/", index)
	fmt.Println("Server is up and listening on port 8000.")
	http.ListenAndServe(":8000", nil)
	
}