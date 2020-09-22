package main

import (
	"log"
	"net/http"
	"whiskey/models"
)

func root(w http.ResponseWriter, r *http.Request) {
	var (
		user models.User
		err  error
	)

	user.Username, err = r.Body
	if err != nil {
		log.Fatal(err)
	}

	w.Write(rep)
}

func main() {
	http.HandleFunc("/test", root)
	http.ListenAndServe("127.0.0.1:8090", nil)
}
