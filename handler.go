package main

import (
	"fmt"
	"net/http"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func userInfoHandler(w http.ResponseWriter, r *http.Request) {
	sub := GetSubject(r)
	if sub == nil {
		fmt.Fprintln(w, "unauthenticated")
	} else {
		fmt.Fprintf(w, "user: %s\n", sub.Name)
		fmt.Fprintf(w, "groups: %v\n", sub.Groups)
	}
}
