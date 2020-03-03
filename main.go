package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

		fmt.Fprintf(w, err.Error())
		status := http.StatusInternalServerError
		http.Error(w, http.StatusText(status), status)
		return
	}
	acc := r.Header.Get("Accept")
	if acc != "" {
		w.Header().Set("Content-Type", acc)
	}

	fmt.Fprintf(w, "%s", body)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
