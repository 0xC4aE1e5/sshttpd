package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time := time.Now().Format(time.RFC1123)
		fmt.Println("[" + time + "] " + r.Method + " " + r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Args[1]
	http.Handle("/", http.FileServer(http.Dir(os.Args[2])))
	fmt.Println("Server started on port " + port)
	err := http.ListenAndServe(":"+port, logger(http.DefaultServeMux))
	if err != nil {
		fmt.Println(err)
	}

}
