package main

import(
	"fmt"    
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hi Ted.")
}

func main(){

	http.HandleFunc("/", get)

	fmt.Println("Starting Server...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error")
	}
}